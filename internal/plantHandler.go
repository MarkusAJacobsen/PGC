package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"pgc/internal/pkg"
)

func plantHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addPlant(w, r)
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

func addPlant(w http.ResponseWriter, r *http.Request) {
	plant := pkg.Plant{}
	pkg.GetPostData(r.Body, &plant, w)

	db := Neo4jPG{}
	if err := db.Connect(); err != nil {
		logrus.Info(err)
		return
	}

	encPlant := CreatePlant(plant)
	if err := db.Create(CreatePlantCypher, encPlant); err != nil {
		fmt.Print(err)
		return
	}
	defer db.Driver.Close()
}
