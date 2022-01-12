package controllers

import (
	"awesomeProject/src/models"
	"awesomeProject/src/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, models.HttpError{Error: "Wrong user structure", Code: 500})
		return
	}
	insertedId, err := services.CreateNewUser(user)
	if err != nil {
		c.JSON(400, models.HttpError{Error: "User could not be added to database", Code: 500})
		return
	}
	c.JSON(200, gin.H{"data": user, "insertedId": insertedId.InsertedID})
}

func GetUserProfile(c *gin.Context) {
	user, err := services.GetUserByEmail(c.GetHeader("email"))
	if err != nil {
		c.JSON(404, models.HttpError{Error: "Can't find user", Code: 404})
		return
	}
	user.Password = ""
	c.JSON(200, user)
}
