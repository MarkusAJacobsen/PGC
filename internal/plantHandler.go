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
}

func addPlant(plant pkg.Plant) {}
