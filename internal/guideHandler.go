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

	idGen := GetIdGenerator()
	g.Id = idGen.GenId(nil)

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
		s.Id = idGen.GenId(nil)
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

	var chTitlesConverted []string
	chapterTitles, ok := rec.Get("chapterTitles")
	if !ok {
		return nil, errors.New("Could not find key 'chapterTitles' in Record")
	} else {
		chTitlesConverted = make([]string, len(chapterTitles.([]interface{})))
		for key, chT := range chapterTitles.([]interface{}) {
			_, ok = chT.(string)
			if ok {
				chTitlesConverted[key] = chT.(string)
			}
		}
	}

	stages, ok := rec.Get("stages")
	if !ok {
		return nil, errors.New("Could not find key 'stages' in Record")
	}

	stageRelations, ok := rec.Get("containsStageRel")
	if !ok {
		return nil, errors.New("Could not find key 'containsStageRel' in Record")
	}

	s := getStages(stages, stageRelations.([]interface{}))

	return pkg.Guide{
		Id:            id.(string),
		Title:         title.(string),
		ChapterTitles: chTitlesConverted,
		Stages:        s,
	}, nil
}

func getStages(stages interface{}, stageRelations []interface{}) []pkg.Stage {
	pageNums := make([]int64, len(stageRelations))
	chapterNums := make([]int64, len(stageRelations))
	filterStrings := make([]string, len(stageRelations))
	for i, rel := range stageRelations {
		props := rel.(neo4j.Relationship).Props()
		pageNums[i] = props["pageNr"].(int64)
		chapterNums[i] = props["chapterNr"].(int64)
		filterStrings[i] = props["filter"].(string)
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
			Id:        raw["id"].(string),
			Title:     raw["title"].(string),
			Text:      raw["text"].(string),
			PageNr:    pageNums[i],
			ChapterNr: chapterNums[i],
			Filter:    filterStrings[i],
			Images:    images,
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
