package internal

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
)

type INeo4jPG interface {
	Connect() (err error)
}

type Neo4jPG struct {
	Driver  neo4j.Driver
	session neo4j.Session
}

func (n *Neo4jPG) Create(cypher string, obj map[string]interface{}) (err error) {
	if n.session, err = n.Driver.Session(neo4j.AccessModeWrite); err != nil {
		logrus.Infoln("Error thrown in session")
		return
	}
	defer n.session.Close()

	if _, err = n.session.Run(cypher, obj); err != nil {
		logrus.Info("error", err)
		return
	}

	return
}

func (n *Neo4jPG) Read() {}

func (n *Neo4jPG) Update() {}

func (n *Neo4jPG) Delete() {}

func (n *Neo4jPG) Connect() (err error) {
	if n.Driver, err = neo4j.NewDriver("bolt://neo4j:testing@neo4j:7687", neo4j.BasicAuth("neo4j", "password", "")); err != nil {
		logrus.Error("Error thrown in driver")
		return
	}

	return
}
