package main

import (
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

var (
	driver  neo4j.Driver
	session neo4j.Session
	result  neo4j.Result
	err     error
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Panic("Port not sat")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handleRoutes)

	if err := http.ListenAndServe("0.0.0.0:"+port, r); err != nil {
		logrus.Error(err)
	}
}

func tryNeo4j() error {
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", "")); err != nil {
		return err // handle error
	}
	// handle driver lifetime based on your application lifetime requirements
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer session.Close()

	result, err = session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
		"id":   1,
		"name": "Item 1",
	})
	if err != nil {
		return err // handle error
	}

	for result.Next() {
		logrus.Info("Created Item with Id = '%d' and Name = '%s'\n", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
	}
	if err = result.Err(); err != nil {
		return err // handle error
	}

	return nil
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
	if err := tryNeo4j(); err != nil {
		logrus.Error(err)
	}
}
