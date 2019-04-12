package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
	"net/http"
	"pgc/internal/pkg"
)

func projectHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		res, err := getUserProjects(r)
		if err != nil {
			WriteServerError(w, err)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			WriteServerError(w, err)
		}

		break
	case http.MethodPost:
		if err := addProject(w, r); err != nil {
			WriteServerError(w, err)
		}
		break
	}
}

func addProject(w http.ResponseWriter, r *http.Request) (err error) {
	pl := pkg.ProjectLink{}
	pkg.GetPostData(r.Body, &pl, w)

	logrus.Infoln("%s", pl)

	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return err
	}

	if err = db.CreateSession(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer db.session.Close()

	encProject := CreateProject(pl.Project)
	if err = db.Do(CreateProjectCypher, encProject); err != nil {
		return err
	}

	encLink := CreateProjectRelation(pl)
	if err = db.Do(LinkProjectCypher, encLink); err != nil {
		return err
	}

	defer db.Driver.Close()

	return nil
}

func getUserProjects(r *http.Request) (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	vars := mux.Vars(r)
	idToken := vars["idToken"]
	logrus.Infoln(idToken)
	params := map[string]interface{}{"idToken": idToken}
	res, err = db.Read(GetProjectsCypher, params)
	if err != nil {
		return nil, err
	}
	logrus.Infoln("%s", res)
	return res, err
}
