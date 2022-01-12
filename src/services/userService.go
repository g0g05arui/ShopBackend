package services

import (
	"awesomeProject/src/database"
	"awesomeProject/src/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func getUsersCollection() *mongo.Collection {
	database := database.DataBase
	return database.Collection("Users")
}

func CreateNewUser(user models.User) (*mongo.InsertOneResult, error) {
	usersCollection := getUsersCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return usersCollection.InsertOne(ctx, user)
}

func GetUserByEmail(email string) (models.User, error) {
	usersCollection := getUsersCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	if err := usersCollection.FindOne(ctx, gin.H{"email": email}).Decode(&user); err != nil {
		return models.User{}, errors.New("no user with this email")
	}
	return user, nil
}
