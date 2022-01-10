package controllers

import (
	"awesomeProject/src/models"
	"awesomeProject/src/services"
	"fmt"
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
	fmt.Println(user.Email, user, insertedId)
}
