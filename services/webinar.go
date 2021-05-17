package services

import (
	"proximity/config"
	"proximity/repo"
)

// WebinarInterface ...
type WebinarInterface interface {
}

//NewWebinarService ..
func NewWebinarService(conf config.IConfig, userDb repo.UserRepoInterface) WebinarInterface {
	return &Webinar{config: conf, user: userDb}
}

// Upload ..
func (i *Webinar) Upload(userName string) (err error) {
	/**
	verify metadata
	upload to blob storage
	**/
	return nil
}

// User ..
type Webinar struct {
	config config.IConfig
	user   repo.UserRepoInterface
}
