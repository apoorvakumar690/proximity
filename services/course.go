package services

import (
	"proximity/config"
	"proximity/models"
	"proximity/repo"

	"gopkg.in/mgo.v2/bson"
)

// CourseInterface ...
type CourseInterface interface {
	GetAllBadWords(skip int, limit int) ([]string, error)
}

//NewCourseService ..
func NewCourseService(conf config.IConfig, userDb repo.CourseRepoInterface) CourseInterface {
	return &Course{config: conf, course: userDb}
}

// GetRole .. get role of a user
func (i *Course) GetCourses(skip int, limit int) (courses []models.CourseMaster, err error) {
	crse, err := i.course.GetAllCourseWithId(bson.M{""}, skip, limit)
	if err != nil {
		return nil, err
	}
	return crse, nil
}

// Course ..
type Course struct {
	config config.IConfig
	course repo.CourseRepoInterface
}
