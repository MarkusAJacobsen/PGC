package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/http"
	"pgc/internal/pkg"
)

func plantHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := addPlant(w, r); err != nil {
			WriteServerError(w, err)
		}
		break
	case http.MethodGet:
		vars := mux.Vars(r)
		var res interface{}
		var err error
		if vars["pId"] != "" {
			res, err = fetchPlant(vars["pId"])
		} else {
			res, err = fetchPlants()
		}

		if err != nil {
			WriteServerError(w, err)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			WriteServerError(w, err)
		}
		break
	case http.MethodDelete:
		if err := deletePlant(r); err != nil {
			WriteServerError(w, err)
		}
		break
	}
}

func plantBatchHandle(w http.ResponseWriter, r *http.Request) {
	var plants []pkg.Plant
	pkg.GetPostData(r.Body, &plants, w)

	var encPlants []map[string]interface{}
	for _, plant := range plants {
		encPlant := CreatePlant(plant)
		encPlants = append(encPlants, encPlant)
	}

	db := Neo4jPG{}
	if err := db.Connect(); err != nil {
		WriteServerError(w, err)
		return
	}

	for i, encPlant := range encPlants {
		if err := db.Create(CreatePlantCypher, encPlant); err != nil {
			WriteServerError(w, err)
			return
		}

		encFamily := CreateFamily(plants[i])
		if err := db.Do(CreatePlantFamilyCypher, encFamily); err != nil {
			WriteServerError(w, err)
			return
		}

		encPlantRelation := CreatePlantRelation(plants[i])
		if err := db.Do(LinkPlantAndFamilyCypher, encPlantRelation); err != nil {
			WriteServerError(w, err)
			return
		}
	}

	defer db.Driver.Close()
}

func addPlant(w http.ResponseWriter, r *http.Request) (err error) {
	plant := pkg.Plant{}
	pkg.GetPostData(r.Body, &plant, w)

	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return err
	}

	if err = db.CreateSession(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer db.Session.Close()

	encPlant := CreatePlant(plant)
	if err = db.Do(CreatePlantCypher, encPlant); err != nil {
		return err
	}

	encFamily := CreateFamily(plant)
	if err = db.Do(CreatePlantFamilyCypher, encFamily); err != nil {
		return err
	}

	encPlantRelation := CreatePlantRelation(plant)
	if err = db.Do(LinkPlantAndFamilyCypher, encPlantRelation); err != nil {
		return err
	}
	defer db.Driver.Close()

	return nil
}

func deletePlant(r *http.Request) (err error) {
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
	pId := vars["pId"]
	param := map[string]interface{}{"id": pId}
	if err = db.Do(DeletePlantCypher, param); err != nil {
		return err
	}

	return err
}

func fetchPlants() (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	res, err = db.Read(GetAllPlantsCypher, nil)
	if err != nil {
		return nil, err
	}

	return res, err
}

func fetchPlant(pId string) (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	param := map[string]interface{}{"pId": pId}
	res, err = db.Read(GetPlantCypher, param)
	if err != nil {
		return nil, err
	}

	return res, err
}
