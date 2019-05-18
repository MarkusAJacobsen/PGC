package internal

import (
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sirupsen/logrus"
	"os"
	"pgc/internal/pkg"
	"time"
)

type INeo4jPG interface {
	Connect() (err error)
	CreateSession(mode neo4j.AccessMode) (err error)
	InitializeConstraints(constrains []string) (err error)
}

type Neo4jPG struct {
	Driver  neo4j.Driver
	Session neo4j.Session
	Cred Neo4jPGCredentials
}

type Neo4jPGCredentials struct {
	URL string
	User string
	Password string
}

var DefaultCred = Neo4jPGCredentials{
	URL: "bolt://neo4j:testing@neo4j:7687",
	User: "neo4j",
	Password: "password",
}

func (n *Neo4jPG) Create(cypher string, obj map[string]interface{}) (err error) {
	if n.Session, err = n.Driver.Session(neo4j.AccessModeWrite); err != nil {
		pkg.ReportError(pgl.ErrorReport{Msg: "Error thrown in Session", Err: err.Error()})
		return err
	}
	defer n.Session.Close()

	var result neo4j.Result
	_, err = n.Session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, obj); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

// Do - Perform Cypher queries on DB, does not close the Session, meaning that you first
// have to create a Session using CreateSession. Subsequently you can perform multiple
// queries using the same Session. CAUTION you have to Close the Session in calling function
func (n *Neo4jPG) Do(cypher string, obj map[string]interface{}) (err error) {
	var result neo4j.Result
	_, err = n.Session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, obj); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

// Read executes with Read only permissions on neo4j.
func (n *Neo4jPG) Read(cypher string, params map[string]interface{}, cRHandle func(r neo4j.Record) (interface{}, error)) (res interface{}, err error) {
	if n.Session, err = n.Driver.Session(neo4j.AccessModeRead); err != nil {
		pkg.ReportError(pgl.ErrorReport{Msg: "Error thrown in Session", Err: err.Error()})
		return nil, err
	}
	defer n.Session.Close()

	res, err = n.Session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []interface{}
		var result neo4j.Result

		if result, err = tx.Run(cypher, params); err != nil {
			return nil, err
		}

		for result.Next() {
			if cRHandle != nil {
				r, err := cRHandle(result.Record())
				if err != nil {
					return nil, err
				}
				list = append(list, r)
			} else {
				props := result.Record().GetByIndex(0).(neo4j.Node).Props()
				list = append(list, props)
			}
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
	n.fetchCredentials()
	n.Driver, err = neo4j.NewDriver(n.Cred.URL, neo4j.BasicAuth(n.Cred.User, n.Cred.Password, ""), func(config *neo4j.Config) {
		config.SocketConnectTimeout = 15 * time.Second
		config.MaxTransactionRetryTime = 15 * time.Second
	})

	if err != nil {
		return
	}

	return
}

func (n *Neo4jPG) CreateSession(mode neo4j.AccessMode) (err error) {
	if n.Session, err = n.Driver.Session(mode); err != nil {
		return err
	}
	return
}

func (n *Neo4jPG) InitializeConstraints(constrains []string) (err error) {
	n.Connect()
	n.CreateSession(neo4j.AccessModeWrite)
	defer n.Session.Close()
	defer n.Driver.Close()

	for _, constraint := range constrains {
		if n.Do(constraint, nil); err != nil {
			return err
		}
	}

	return nil
}

func (n *Neo4jPG) fetchCredentials()  {
	gUrl := os.Getenv("GRAPHENEDB_BOLT_URL")
	gUser := os.Getenv("GRAPHENEDB_BOLT_USER")
	gPassword := os.Getenv("GRAPHENEDB_BOLT_PASSWORD")

	if gUrl == "" && gUser == "" && gPassword == "" {
		logrus.Infoln("Using default credentials")
		n.Cred = DefaultCred
	} else {
		logrus.Infoln("Using GrapheneDB")
		n.Cred.URL = gUrl
		n.Cred.User = gUser
		n.Cred.Password = gPassword
	}
}
