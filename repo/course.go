package repo

import (
	"proximity/config"
	"proximity/models"
	"proximity/pkg/clients/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const courseColNmae = "course"

// CourseRepoInterface ..
type CourseRepoInterface interface {
	Insert(u models.CourseMaster) error
	GetAllCourseWithId(filter primitive.M, skip int, limit int) ([]models.CourseMaster, error)
	UpdateOne(filter primitive.M, update primitive.M) error
	DeleteOne(filter primitive.M) error
}

// NewCourseRepo Create's an instance of a Webinr Repository
func NewCourseRepo(conf config.IConfig, dbInstances *db.Instances) CourseRepoInterface {
	return &CourseRepo{config: conf, db: dbInstances.MongoDB}
}

// BadWordsRepo Contains methods to action on the User Repository
type CourseRepo struct {
	config config.IConfig
	db     db.MongoDBI
}

// Insert .. Insert new Course
func (ur *CourseRepo) Insert(u models.CourseMaster) error {
	_, err := ur.db.InsertOne(courseColNmae, u)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne ...
func (ur *CourseRepo) UpdateOne(filter primitive.M, update primitive.M) error {
	_, err := ur.db.UpdateOne(courseColNmae, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOne ...
func (ur *CourseRepo) DeleteOne(filter primitive.M) error {
	_, err := ur.db.DeleteOne(courseColNmae, filter)
	if err != nil {
		return err
	}
	return nil
}

// GetAllWebinarWithId ...
func (ur *CourseRepo) GetAllCourseWithId(filter primitive.M, skip int, limit int) ([]models.CourseMaster, error) {
	var ws []models.CourseMaster
	_, err := ur.db.FindManyAndPaginate(courseColNmae, filter, skip, limit, "updatedAt", &ws)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
