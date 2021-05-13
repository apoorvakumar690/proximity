package db

import (
	"proximity/config"
)

// Instances ... contains the interface layer of the different dbs
type Instances struct {
	IdyaMongoDB MongoDBI
}

// NewInstance creates an instance of initialized DBInstances
func NewInstance(conf config.IConfig) (*Instances, error) {
	dbInstances := &Instances{}

	idyaMongoDB, err := initIdyaMongoDB(conf)
	if err != nil {
		return nil, err
	}

	// Sets db instance
	dbInstances.IdyaMongoDB = idyaMongoDB

	return dbInstances, nil
}

// Simulates the initialization of a db connection
//initRnRMongoDB ..
func initIdyaMongoDB(config config.IConfig) (MongoDBI, error) {
	return NewMongoDB(config)
}
