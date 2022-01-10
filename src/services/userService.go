package services

import (
	"awesomeProject/src/database"
	"awesomeProject/src/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func getUsersCollection() *mongo.Collection {
	database := database.DataBase
	return database.Collection("Users")
}

/*
	TODO check for user validity email,password,phoneNo,name and firstName are not empty
*/

func CreateNewUser(user models.User) (*mongo.InsertOneResult, error) {
	usersCollection := getUsersCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return usersCollection.InsertOne(ctx, user)
}
