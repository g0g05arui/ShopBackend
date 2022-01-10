package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mongoDBUri = `mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false`
var MongoClient *mongo.Client
var DataBase *mongo.Database

func Setup() {
	clientOptions := options.Client().ApplyURI(mongoDBUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient, _ = mongo.Connect(ctx, clientOptions)
	DataBase = MongoClient.Database("Shop")
}
