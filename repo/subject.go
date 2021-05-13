package repo

import (
	"proximity/config"
	"proximity/models"
	"proximity/pkg/clients/db"

	"go.mongodb.org/mongo-driver/bson"
)

const userColNmae = "badwords"

// UserRepoInterface ..
type UserRepoInterface interface {
	Insert(u models.User) error
	GetOne(userName string) (*models.User, error)
}

// NewUserRepo Create's an instance of a User Repository
func NewUserRepo(conf config.IConfig, dbInstances *db.Instances) UserRepoInterface {
	return &UserRepo{config: conf, db: dbInstances.RnrMongoDB}
}

// BadWordsRepo Contains methods to action on the User Repository
type UserRepo struct {
	config config.IConfig
	db     db.MongoDBI
}

// Insert .. Insert new User
func (ur *UserRepo) Insert(u models.User) error {
	_, err := ur.db.InsertOne(userColNmae, u)

	if err != nil {
		return err
	}
	return nil
}

// GetOne ..
func (ur *UserRepo) GetOne(userName string) (*models.User, error) {
	w := ur.db.FindOne(userColNmae, bson.M{"userName": userName})

	var s models.User
	if err := w.Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}
