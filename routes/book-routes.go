package routes

import (
	"test/controllers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	bookGroup := router.Group("/books").Use(middlewares.AuthMiddleware())
	{
		bookGroup.GET("/", controllers.GetAllBooks)
		bookGroup.POST("/", controllers.CreateBook)
	}
}
