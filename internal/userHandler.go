package internal

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/http"
	"pgc/internal/pkg"
)

func userHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := addUser(w, r); err != nil {
			WriteServerError(w, err)
		}
		break
	case http.MethodPut:
		if err := updateUser(w, r); err != nil {
			WriteServerError(w, err)
		}
		break
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) (err error) {
	u := pkg.User{}
	pkg.GetPostData(r.Body, &u, w)

	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return err
	}

	if err = db.CreateSession(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer db.session.Close()

	encU := CreateUser(u)
	if err = db.Do(CreateUserCypher, encU); err != nil {
		return err
	}

	if u.Area != "" {
		encArea := CreateArea(u)
		if err = db.Do(CreateAreaCypher, encArea); err != nil {
			return err
		}

		encUserAreaRelation := CreateUserAreaRelation(u)
		if err = db.Do(LinkUserAndAreaCypher, encUserAreaRelation); err != nil {
			return err
		}
	}

	defer db.Driver.Close()

	return nil
}

func addUser(w http.ResponseWriter, r *http.Request) (err error) {
	u := pkg.User{}
	pkg.GetPostData(r.Body, &u, w)

	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return err
	}

	encU := CreateUser(u)
	if err = db.Create(CreateUserCypher, encU); err != nil {
		return err
	}

	defer db.Driver.Close()

	return nil
}
