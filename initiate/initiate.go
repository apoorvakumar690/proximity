package initiate

import (
	"proximity/apis"
	"proximity/config"
	"proximity/pkg/clients/db"
	httpPkg "proximity/pkg/clients/http"
	log "proximity/pkg/utils/logger"
	"proximity/shared"
)

// Initialize will initialize all the dependencies and the servers.
// Dependencies include config, external connections(grpc, http)
func Initialize() error {
	// Initializes logger
	log.InitLogger()

	env, err := Env()
	if err != nil {
		return err
	}

	log.Info("Enviroment: " + env)

	// Sets apm env
	SetApmEnv(env)

	// Initializes APM based on environment
	// if env != config.ENVDevelopment && env != config.ENVDocker && env != config.ENVSit {
	// 	// remove sit when building an actual app
	// 	apm.Initialize()
	// }

	// Gets config
	conf, err := config.NewConfig(env)
	if err != nil {
		return err
	}

	// Initializes the DB connections
	dbInstances, err := db.NewInstance(conf)
	if err != nil {
		return err
	}

	// loads all common dependencies
	dependencies := shared.Deps{
		Config:        conf,
		Database:      dbInstances,
		HTTPRequester: httpPkg.NewRequest(),
	}

	// Initializes servers
	err = apis.InitServers(&dependencies)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
