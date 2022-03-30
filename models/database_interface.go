package models

import "go.mongodb.org/mongo-driver/mongo"

type Database interface {
	FindOne(filter interface{}, result interface{}) error
	FindAll(filter interface{}, result interface{}) error
	UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	InsertOne(document interface{}) (*mongo.InsertOneResult, error)
}
