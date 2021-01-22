package routers

import (
	"test/models/services"

	"github.com/gin-gonic/gin"
)

//InitRoute using Gin
func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")
	client := router.Group("/")
	client.DELETE("/user/:id", services.DeleteUser)
	client.GET("/user/:id", services.GetUser)
	client.PUT("/user/:id", services.UpdateUser)
	client.GET("/user", services.GetAllUser)
	client.POST("/user", services.Create)

	return router
}
