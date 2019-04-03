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

// Do - Perform Cypher queries on DB, does not close the session, meaning that you first
// have to create a session using CreateSession. Subsequently you can perform multiple
// queries using the same session. CAUTION you have to Close the session in calling function
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

func (n *Neo4jPG) Read(cypher string) (res interface{}, err error) {
	if n.session, err = n.Driver.Session(neo4j.AccessModeRead); err != nil {
		logrus.Infoln("Error thrown in session")
		return nil, err
	}
	defer n.session.Close()

	res, err = n.session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []map[string]interface{}
		var result neo4j.Result

		if result, err = tx.Run(cypher, nil); err != nil {
			return nil, err
		}

		for result.Next() {
			list = append(list, result.Record().GetByIndex(0).(neo4j.Node).Props())
		}

		if err = result.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})

	if err != nil {
		return nil, err
	}

	return res, err
}

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

func (n *Neo4jPG) InitializeConstraints(constrains []string) (err error) {
	n.Connect()
	n.CreateSession(neo4j.AccessModeWrite)
	defer n.session.Close()
	defer n.Driver.Close()

	for _, constraint := range constrains {
		if n.Do(constraint, nil); err != nil {
			return err
		}
	}

	return nil
}
