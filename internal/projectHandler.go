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
	case http.MethodDelete:
		if err := deleteProject(r); err != nil {
			WriteServerError(w, err)
		}
		break
	}
}

func userProjectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		res, err := getUserProject(r)
		if err != nil {
			WriteServerError(w, err)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
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
	defer db.Driver.Close()

	if err = db.CreateSession(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer db.Session.Close()

	encProject := CreateProject(pl.Project)
	if err = db.Do(CreateProjectCypher, encProject); err != nil {
		return err
	}

	encLink := CreateProjectRelation(pl)
	if err = db.Do(LinkProjectCypher, encLink); err != nil {
		return err
	}

	return nil
}

func getUserProjects(r *http.Request) (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	vars := mux.Vars(r)
	idToken := vars["uIdToken"]
	params := map[string]interface{}{"idToken": idToken}
	res, err = db.Read(GetProjectsCypher, params)
	if err != nil {
		return nil, err
	}
	logrus.Infof("%s", res)
	return res, err
}

func getUserProject(r *http.Request) (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	vars := mux.Vars(r)
	idToken := vars["uIdToken"]
	projectId := vars["pId"]
	params := map[string]interface{}{"idToken": idToken, "pId": projectId}
	res, err = db.Read(GetProjectCypher, params)
	if err != nil {
		return nil, err
	}
	logrus.Infof("%s", res)
	return res, err
}

func deleteProject(r *http.Request) (err error) {
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
	projectId := vars["pId"]
	params := map[string]interface{}{"id": projectId}
	if err = db.Do(DeleteProjectCypher, params); err != nil {
		return err
	}

	return err
}
