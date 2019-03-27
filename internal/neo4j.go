package internal

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
	"time"
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
		return err
	}
	defer n.session.Close()

	var result neo4j.Result
	_, err = n.session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, obj); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

func (n *Neo4jPG) Do(cypher string, obj map[string]interface{}) (err error) {
	var result neo4j.Result
	_, err = n.session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, obj); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

func (n *Neo4jPG) Read() {}

func (n *Neo4jPG) Update() {}

func (n *Neo4jPG) Delete() {}

func (n *Neo4jPG) Connect() (err error) {
	n.Driver, err = neo4j.NewDriver("bolt://neo4j:testing@neo4j:7687", neo4j.BasicAuth("neo4j", "password", ""), func(config *neo4j.Config) {
		config.SocketConnectTimeout = 15 * time.Second
		config.MaxTransactionRetryTime = 15 * time.Second
	})

	if err != nil {
		return
	}

	return
}

func (n *Neo4jPG) CreateSession(mode neo4j.AccessMode) (err error) {
	if n.session, err = n.Driver.Session(mode); err != nil {
		return err
	}
	return
}
