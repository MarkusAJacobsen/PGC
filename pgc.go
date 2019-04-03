package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"pgc/internal"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Panic("Port not sat")
	}

	r := internal.SetUpRouter()

	printStartUpMsg(port)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	db := internal.Neo4jPG{}
	if err := db.InitializeConstraints([]string{internal.UserConstraintCypher}); err != nil {
		logrus.Errorln("Constraints could not be established", err.Error())
	}

	logrus.Fatal(srv.ListenAndServe())
}

func printStartUpMsg(port string) {
	logrus.Infof("Starting up PGC on port %s", port)
}
