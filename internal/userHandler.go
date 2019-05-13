package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/http"
	"pgc/internal/pkg"
)

func userHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		res, err := getUser(r)
		if err != nil {
			WriteServerError(w, err)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			WriteServerError(w, err)
		}

		break
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
	case http.MethodDelete:
		if err := deleteUser(r); err != nil {
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
	defer db.Session.Close()

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

func getUser(r *http.Request) (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	vars := mux.Vars(r)
	idToken := vars["uIdToken"]
	param := map[string]interface{}{"idToken": idToken}
	res, err = db.Read(GetUserCypher, param)
	if err != nil {
		return nil, err
	}

	return res, err
}

func deleteUser(r *http.Request) (err error) {
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
	idToken := vars["uIdToken"]
	param := map[string]interface{}{"idToken": idToken}
	if err = db.Do(DeleteUserCypher, param); err != nil {
		return err
	}

	return err
}
