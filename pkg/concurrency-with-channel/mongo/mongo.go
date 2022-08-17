package mongo

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository interface {
	ReadOne(collectionName string, read ReadOne) (response []byte, err error)
	ReadMany(collectionName string, readMany ReadMany) (response []byte, err error)
	UpdateOne(collectionName string, update UpdateOne) (err error)
	UpdateMany(collectionName string, updateMany UpdateMany) (err error)
	InsertOne(collectionName string, insert InsertOne) (err error)
	InsertMany(collectionName string, insert InsertMany) (err error)
	Aggregate(collectionName string, aggregate Aggregate) (results []byte, err error)
}

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) MongoRepository {
	return &mongoRepository{db}
}

func (r *mongoRepository) ReadOne(collectionName string, readOne ReadOne) (response []byte, err error) {

	var result bson.M

	if err = r.db.Collection(collectionName).FindOne(context.TODO(), readOne.filter, readOne.opts).Decode(&result); err != nil {
		return nil, errors.New("No document")
	}

	if response, err = json.Marshal(result); err != nil {
		return nil, err
	}

	return response, err
}

func (r *mongoRepository) UpdateOne(collectionName string, updateOne UpdateOne) (err error) {
	var res *mongo.UpdateResult

	if res, err = r.db.Collection(collectionName).UpdateOne(context.TODO(), updateOne.filter, updateOne.update, updateOne.opts); err != nil {
		return
	}

	if res.MatchedCount == 0 {
		return errors.New("MongoDb UpdateOne Failed")
	}
	return
}

func (r *mongoRepository) InsertOne(collectionName string, insertOne InsertOne) (err error) {
	_, err = r.db.Collection(collectionName).InsertOne(context.TODO(), insertOne.data, insertOne.opts)
	return
}

func (r *mongoRepository) ReadMany(collectionName string, readMany ReadMany) (result []byte, err error) {
	var cursor *mongo.Cursor
	var rows []bson.M
	var offset int64 = 0
	var limit int64 = 100

	if readMany.opts == nil {
		readMany.opts = options.Find()
	}

	for true {
		var limitedRows []bson.M
		readMany.opts.SetSkip(offset).SetLimit(limit)

		if cursor, err = r.db.Collection(collectionName).Find(context.TODO(), readMany.filter, readMany.opts); err != nil {
			return
		}

		if err = cursor.All(context.TODO(), &limitedRows); err != nil {
			return
		}

		if len(limitedRows) == 0 {
			break
		}

		offset += limit
		rows = append(rows, limitedRows...)
	}

	if result, err = json.Marshal(rows); err != nil {
		return
	}

	return
}

func (r *mongoRepository) UpdateMany(collectionName string, updateMany UpdateMany) (err error) {
	var res *mongo.UpdateResult

	if res, err = r.db.Collection(collectionName).UpdateMany(context.TODO(), updateMany.filter, updateMany.update, updateMany.opts); err != nil {
		return
	}

	if res.MatchedCount == 0 {
		return errors.New("MongoDb UpdateOne Failed")
	}
	return
}

func (r *mongoRepository) InsertMany(collectionName string, insertMany InsertMany) (err error) {
	_, err = r.db.Collection(collectionName).InsertMany(context.TODO(), insertMany.data, insertMany.opts)
	return
}

func (r *mongoRepository) Aggregate(collectionName string, aggregate Aggregate) (results []byte, err error) {
	var cursor *mongo.Cursor

	if cursor, err = r.db.Collection(collectionName).Aggregate(context.TODO(), aggregate.pipeline, aggregate.opts); err != nil {
		return
	}

	var rows []bson.M
	if err = cursor.All(context.TODO(), &rows); err != nil {
		return
	}
	return json.Marshal(rows)
}
