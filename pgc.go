package main

import (
	"github.com/MarkusAJacobsen/PGC/internal"
	"github.com/MarkusAJacobsen/PGC/internal/pkg"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const DefaultLogURL = "http://172.19.0.3:6113/report"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Panic("Port not sat")
	}

	logURL := os.Getenv("LOG_URL")
	if logURL == "" {
		logrus.Infoln("Using default log URL")
		logURL = DefaultLogURL
	}

	s := pkg.Server{
		LoggerURL: logURL,
	}

	r := internal.SetUpRouter()
	r.Use(s.TrafficMiddleware)

	printStartUpMsg(port)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Fatal(srv.ListenAndServe())
}

func printStartUpMsg(port string) {
	logrus.Infof("Starting up PGC on port %s", port)
}
