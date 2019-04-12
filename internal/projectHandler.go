package internal

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
	"net/http"
	"pgc/internal/pkg"
)

func projectHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
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
