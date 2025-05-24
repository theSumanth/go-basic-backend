package main

import (
	"example.com/go-basic-backend/db"
	"example.com/go-basic-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
