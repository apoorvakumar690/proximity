package services

import (
	"proximity/config"
	"proximity/repo"
)

// UserInterface ...
type UserInterface interface {
	GetRole(userName string) (role string, err error)
}

//NewUserService ..
func NewUserService(conf config.IConfig, userDb repo.UserRepoInterface) UserInterface {
	return &User{config: conf, user: userDb}
}

// GetRole .. get role of a user
func (i *User) GetRole(userName string) (role string, err error) {
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
