package services

import (
	"proximity/config"
	"proximity/repo"
)

// UserInterface ...
type UserInterface interface {
	GetAllBadWords(skip int, limit int) ([]string, error)
}

//NewUserService ..
func NewUserService(conf config.IConfig, userDb repo.UserRepoInterface) CmsInterface {
	return &Cms{config: conf, user: userDb}
}

// GetRole .. get role of a user
func (i *Cms) GetRole(userName string) (role string, err error) {
	wds, err := i.user.GetOne(userName)
	if err != nil {
		return nil, err
	}
	for _, w := range wds {
		words = append(words, w.Word)
	}
	return words, nil
}

// User ..
type User struct {
	config config.IConfig
	user   repo.UserRepoInterface
}
