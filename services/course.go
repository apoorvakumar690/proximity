package services

import (
	"proximity/config"
	"proximity/repo"
)

// CourseInterface ...
type CourseInterface interface {
	GetAllBadWords(skip int, limit int) ([]string, error)
}

//NewCourseService ..
func NewCourseService(conf config.IConfig, userDb repo.UserRepoInterface) CmsInterface {
	return &Cms{config: conf, user: userDb}
}

// GetRole .. get role of a user
func (i *Course) GetCourses() (courses [], err error) {
	wds, err := i.course.GetMany(skip string)
	if err != nil {
		return nil, err
	}
	return rold, nil
}

// Course ..
type Course struct {
	config config.IConfig
	user   repo.UserRepoInterface
}
