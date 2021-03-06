package db

import (
	"proximity/config"
)

// Instances ... contains the interface layer of the different dbs
type Instances struct {
	MongoDB MongoDBI
}

// NewInstance creates an instance of initialized DBInstances
func NewInstance(conf config.IConfig) (*Instances, error) {
	dbInstances := &Instances{}

	mongoDB, err := initMongoDB(conf)
	if err != nil {
		return nil, err
	}

	// Sets db instance
	dbInstances.MongoDB = mongoDB

	return dbInstances, nil
}

// Simulates the initialization of a db connection
//initRnRMongoDB ..
func initMongoDB(config config.IConfig) (MongoDBI, error) {
	return NewMongoDB(config)
}
