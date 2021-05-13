package repo

import (
	"proximity/config"
	"proximity/pkg/clients/db"
)

// UserRepoInterface ...
type UserRepoInterface interface {
}

// NewUserRepo Create's an instance of a User Repository
func NewUserRepo(conf config.IConfig, dbInstances *db.Instances) UserRepoInterface {
	return &UserRepo{config: conf}
}

// UserRepo Contains methods to action on the User Repository
type UserRepo struct {
	config config.IConfig
}
