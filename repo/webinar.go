package repo

import (
	"proximity/config"
	"proximity/models"
	"proximity/pkg/clients/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const webinarColNmae = "webinar"

// WebinarRepoInterface ..
type WebinarRepoInterface interface {
	Insert(u models.WebinarMaster) error
	GetAllWebinarWithId(filter primitive.M, skip int, limit int) ([]models.WebinarMaster, error)
	UpdateOne(filter primitive.M, update primitive.M) error
	DeleteOne(filter primitive.M) error
}

// NewWebinarRepo Create's an instance of a Webinr Repository
func NewWebinarRepo(conf config.IConfig, dbInstances *db.Instances) WebinarRepoInterface {
	return &WebinarRepo{config: conf, db: dbInstances.MongoDB}
}

// BadWordsRepo Contains methods to action on the User Repository
type WebinarRepo struct {
	config config.IConfig
	db     db.MongoDBI
}

// Insert .. Insert new Webinar
func (ur *WebinarRepo) Insert(u models.WebinarMaster) error {
	_, err := ur.db.InsertOne(webinarColNmae, u)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne ...
func (ur *WebinarRepo) UpdateOne(filter primitive.M, update primitive.M) error {
	_, err := ur.db.UpdateOne(webinarColNmae, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOne ...
func (ur *WebinarRepo) DeleteOne(filter primitive.M) error {
	_, err := ur.db.DeleteOne(webinarColNmae, filter)
	if err != nil {
		return err
	}
	return nil
}

// GetAllWebinarWithId ...
func (ur *WebinarRepo) GetAllWebinarWithId(filter primitive.M, skip int, limit int) ([]models.WebinarMaster, error) {
	var ws []models.WebinarMaster
	_, err := ur.db.FindManyAndPaginate(webinarColNmae, filter, skip, limit, "updatedAt", &ws)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
