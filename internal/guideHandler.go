package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/pkg/errors"
	"net/http"
	"pgc/internal/pkg"
)

func GuideHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		res, err := getGuide(r)
		if err != nil {
			WriteServerError(w, err)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			WriteServerError(w, err)
		}

		break
	case http.MethodPost:
		if err := addGuide(w, r); err != nil {
			WriteServerError(w, err)
		}
		break
	case http.MethodDelete:
		if err := deleteGuide(r); err != nil {
			WriteServerError(w, err)
		}
		break
	}
}

func addGuide(w http.ResponseWriter, r *http.Request) (err error) {
	g := pkg.Guide{}
	pkg.GetPostData(r.Body, &g, w)

	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return err
	}
	defer db.Driver.Close()

	encG := CreateGuide(g)

	if err = db.CreateSession(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer db.Session.Close()

	if err = db.Do(CreateGuideCypher, encG); err != nil {
		return err
	}

	for _, s := range g.Stages {
		encS := CreateStage(s)
		encS["gId"] = g.Id
		if err = db.Do(CreateStageCypher, encS); err != nil {
			return err
		}
	}

	return nil
}

func getGuide(r *http.Request) (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	vars := mux.Vars(r)
	gId := vars["gId"]
	param := map[string]interface{}{"id": gId}
	res, err = db.Read(GetGuideCypher, param, handleGetRecord)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func handleGetRecord(rec neo4j.Record) (res interface{}, err error) {
	id, ok := rec.Get("id")
	if !ok {
		return nil, errors.New("Could not find key 'id' in Record")
	}
	title, ok := rec.Get("title")
	if !ok {
		return nil, errors.New("Could not find key 'title' in Record")
	}
	stages, ok := rec.Get("stages")
	if !ok {
		return nil, errors.New("Could not find key 'stages' in Record")
	}
	pageNumbers, ok := rec.Get("pageNumbers")
	if !ok {
		return nil, errors.New("Could not find key 'pageNumbers' in Record")
	}

	s := getStages(stages, pageNumbers.([]interface{}))

	return pkg.Guide{
		Id:    id.(string),
		Title: title.(string),
		Stages: s,
	}, nil
}

func getStages(stages interface{}, pageNumbers []interface{}) ([]pkg.Stage) {
	pageNrArr := make([]int64, len(pageNumbers))
	for i, pN := range pageNumbers {
		pageNrArr[i] = pN.(neo4j.Relationship).Props()["pageNr"].(int64)
	}

	s := make([]pkg.Stage, len(stages.([]interface{})))
	for i, v := range stages.([]interface{}) {
		raw := v.(neo4j.Node).Props()

		var images []string
		imArr, ok := raw["images"].([]interface{})
		if ok {
			images = make([]string, len(imArr))
			for key, im := range imArr {
				_, ok = im.(string)
				if ok {
					images[key] = im.(string)
				}
			}
		}

		s[i] = pkg.Stage{
			Id: raw["id"].(string),
			Text: raw["text"].(string),
			PageNr: pageNrArr[i],
			Images: images,
		}
	}
	return s
}

func deleteGuide(r *http.Request) (err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return err
	}
	defer db.Driver.Close()

	if err = db.CreateSession(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer db.Session.Close()

	vars := mux.Vars(r)
	id := vars["gId"]
	param := map[string]interface{}{"id": id}
	if err = db.Do(DeleteGuideCypher, param); err != nil {
		return err
	}

	if err = db.Do(DeleteOrphanedStages, nil); err != nil {
		return err
	}

	return err
}