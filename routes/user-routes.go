package routes

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/login", controllers.Login)
		userGroup.POST("/register", controllers.Register)
	}
}
