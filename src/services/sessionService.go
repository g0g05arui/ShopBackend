package services

import (
	"awesomeProject/src/database"
	"awesomeProject/src/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func getSessionsCollection() *mongo.Collection {
	database := database.DataBase
	return database.Collection("Sessions")
}

func CanAuth(email, password string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := getUsersCollection()
	data := userCollection.FindOne(ctx, gin.H{"email": email, "password": password})
	return data.Err() == nil
}

func AddSession(email, password string) (interface{}, error) {
	if CanAuth(email, password) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		sessionsCollection := getSessionsCollection()
		insertedId, _ := sessionsCollection.InsertOne(ctx, models.Session{Id: "0", Email: email, Expires: time.Now().Local().Add(time.Hour * time.Duration(720)), CreatedAt: time.Now().Local()})
		return insertedId.InsertedID, nil
	}
	return nil, errors.New("invalid login credentials")
}

func IsSessionValid(sessionId, userEmail string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	sessionCollection := getSessionsCollection()
	objID, _ := primitive.ObjectIDFromHex(sessionId)
	data := sessionCollection.FindOne(ctx, gin.H{"_id": objID, "email": userEmail})
	return data.Err() == nil
}
