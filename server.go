package main

import (
	"test/db"
	"test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.InitializeRoutes(server)

	db.ConnectDB()

	server.Run(":4000")
}
