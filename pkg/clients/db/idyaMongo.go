package db

import (
	"context"
	"errors"
	"fmt"
	"proximity/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBI ..
type MongoDBI interface {
	InsertOne(string, interface{}) (*mongo.InsertOneResult, error)
	UpdateOne(string, interface{}, interface{}) (*mongo.UpdateResult, error)
	FindOne(string, interface{}) *mongo.SingleResult
	DeleteOne(string, primitive.M) (*mongo.DeleteResult, error)
	InsertMany(string, []interface{}) (*mongo.InsertManyResult, error)
	UpdateMany(string, interface{}, interface{}) (*mongo.UpdateResult, error)
	FindMany(string, interface{}, interface{}) (interface{}, error)
	DeleteMany(string, interface{}) (*mongo.DeleteResult, error)
	FindOneAndUpdate(string, interface{}, interface{}) *mongo.SingleResult
	FindManyAndPaginate(string, interface{}, int, int, string, interface{}) (interface{}, error)
	FindOneAndUpsert(collection string, filter interface{}, update interface{}) *mongo.SingleResult
	Pipe(string, []primitive.M, interface{}) (interface{}, error)
	Count(string, interface{}) (int64, error)
}

// NewMongoDB ..
func NewMongoDB(config config.IConfig) (MongoDBI, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Get().Mongo.URL))
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to mongo " + config.Get().Mongo.URL)
	return &RnRMongoDB{client: client, database: config.Get().Mongo.Database}, err
}

// RnRMongoDB ..
type RnRMongoDB struct {
	client   *mongo.Client
	database string
	err      error
}

//InsertOne function
func (r *RnRMongoDB) InsertOne(collection string, i interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.InsertOne(ctx, i)
}

//UpdateOne function
func (r *RnRMongoDB) UpdateOne(collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.UpdateOne(ctx, filter, update)
}

//FindOne function
func (r *RnRMongoDB) FindOne(collection string, filter interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.FindOne(ctx, filter)
}

//DeleteOne function
func (r *RnRMongoDB) DeleteOne(collection string, i primitive.M) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.DeleteOne(ctx, i)
}

//InsertMany function
func (r *RnRMongoDB) InsertMany(collection string, i []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.InsertMany(ctx, i)
}

//UpdateMany function
func (r *RnRMongoDB) UpdateMany(collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.UpdateMany(ctx, filter, update)
}

//FindMany function
func (r *RnRMongoDB) FindMany(collection string, filter interface{}, result interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	cursor, err := col.Find(ctx, filter)
	defer cursor.Close(ctx)

	err = cursor.All(ctx, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//DeleteMany function
func (r *RnRMongoDB) DeleteMany(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.DeleteMany(ctx, filter)
}

//FindOneAndUpdate function
func (r *RnRMongoDB) FindOneAndUpdate(collection string, filter interface{}, updateDoc interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.M{
		"$set": updateDoc,
	}
	col := r.client.Database(r.database).Collection(collection)
	return col.FindOneAndUpdate(ctx, filter, update)
}

//FindOneAndUpsert function
func (r *RnRMongoDB) FindOneAndUpsert(collection string, filter interface{}, updateDoc interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{
		"$set": updateDoc,
	}
	col := r.client.Database(r.database).Collection(collection)
	return col.FindOneAndUpdate(ctx, filter, update, &opt)
}

//FindManyAndPaginate function
func (r *RnRMongoDB) FindManyAndPaginate(collection string, filter interface{}, skip int, limit int, column string, result interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)

	options := options.Find()

	//sort if only column name passed
	options.SetSort(bson.D{primitive.E{Key: column, Value: -1}})

	options.SetSkip(int64(skip))

	options.SetLimit(int64(limit))

	cursor, err := col.Find(ctx, filter, options)
	defer cursor.Close(ctx)

	err = cursor.All(ctx, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Pipe used for aggregation
func (r *RnRMongoDB) Pipe(collection string, query []primitive.M, res interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	cursor, err := col.Aggregate(ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// FindAllAndSort find all documents with query in sorted order
func (r *RnRMongoDB) FindAllAndSort(collection string, query, project interface{}, sortBy []string, result interface{}) error {
	query, ok := query.(map[string]interface{})
	if !ok {
		return errors.New("Find All and Sort error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)

	sortQuery := buildSortQuery(sortBy)
	cursor, err := col.Find(ctx, query, &options.FindOptions{
		Projection: project,
		Sort:       sortQuery,
	})

	defer cursor.Close(ctx)

	err = cursor.All(ctx, result)

	return err
}

func buildSortQuery(sortBy []string) interface{} {
	sortQuery := map[string]int{}

	for _, key := range sortBy {
		sortQuery[key] = -1
	}
	return sortQuery
}

//Count function
func (r *RnRMongoDB) Count(collection string, filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := r.client.Database(r.database).Collection(collection)
	return col.CountDocuments(ctx, filter)
}
