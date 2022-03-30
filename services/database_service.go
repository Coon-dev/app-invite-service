package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabase struct {
	Collection *mongo.Collection
}

func (mg *MongoDatabase) FindOne(filter interface{}, result interface{}) error {
	return mg.Collection.FindOne(context.Background(), filter).Decode(result)
}

func (mg *MongoDatabase) FindAll(filter interface{}, result interface{}) error {
	cur, err := mg.Collection.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	return cur.All(context.Background(), result)
}

func (mg *MongoDatabase) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return mg.Collection.UpdateOne(context.Background(), filter, update)
}

func (mg *MongoDatabase) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return mg.Collection.InsertOne(context.Background(), document)
}
