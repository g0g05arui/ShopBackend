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

func LoginUser(c *gin.Context) {
	var userLogin models.UserAuth
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(400, models.HttpError{Error: "Wrong user structure", Code: 400})
		return
	}
	if !services.CanAuth(userLogin.Email, userLogin.Password) {
		c.JSON(404, models.HttpError{Error: "Wrong credentials", Code: 404})
		return
	}
	insertedId, err := services.AddSession(userLogin.Email, userLogin.Password)
	if err != nil {
		c.JSON(500, models.HttpError{
			Error: "Internal Error",
			Code:  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"sessionId": insertedId,
	})
}
