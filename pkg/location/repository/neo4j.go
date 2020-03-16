package repository

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/yowmamasita/dctx/pkg/location/domain"
)

type neo4jRepo struct {
	driver neo4j.Driver
}

func New(uri, username, password string, enableTLS bool) domain.LocationService {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), func(c *neo4j.Config) {
		c.Encrypted = enableTLS
	})
	if err != nil {
		log.Fatal("Error creating neo4j driver", err)
	}

	return &neo4jRepo{
		driver,
	}
}

func (r *neo4jRepo) Match() {
	session, err := r.driver.Session(neo4j.AccessModeRead)
	if err != nil {
		log.Fatal("Error creating neo4j read session", err)
	}

	greeting, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("MATCH (n) RETURN n", nil)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		log.Fatal("Match query test failed", err)
	}

	fmt.Println(greeting)
}
