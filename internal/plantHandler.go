package internal

import (
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
	"net/http"
	"pgc/internal/pkg"
)

func plantHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := addPlant(w, r); err != nil {
			logrus.Errorln(err)
		}
		break
	case http.MethodGet:
		res, err := fetchPlants()
		if err != nil {
			logrus.Errorln(err)
			w.WriteHeader(500)
			w.Write([]byte("An error occurred"))
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			logrus.Errorln(err)
			w.WriteHeader(500)
			w.Write([]byte("An error occurred"))

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
		logrus.Info(err)
		return
	}

	for _, encPlant := range encPlants {
		if err := db.Create(CreatePlantCypher, encPlant); err != nil {
			fmt.Print(err)
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
	defer db.session.Close()

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

func fetchPlants() (res interface{}, err error) {
	db := Neo4jPG{}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	defer db.Driver.Close()

	res, err = db.Read(GetAllPlants)
	if err != nil {
		return nil, err
	}

	return res, err
}
