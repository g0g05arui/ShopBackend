package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	application := gin.Default()

	unauthorized := application.Group("/")
	{
		unauthorized.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Server running")
		})
		unauthorized.POST("/user", CreateUser)
		unauthorized.POST("/auth", LoginUser)
	}
	secured := application.Group("/")
	{
		secured.Use(AuthMiddleware)
		secured.GET("/test-secured", func(c *gin.Context) {
			c.String(http.StatusOK, "Security running")
		})
	}
	return application
}
