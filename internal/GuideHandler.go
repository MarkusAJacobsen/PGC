package internal

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/http"
	"pgc/internal/pkg"
)

func GuideHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := addGuide(w, r); err != nil {
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
