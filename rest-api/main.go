package main

import (
	"github.com/Thanasak1412/go-rest-api/db"
	"github.com/Thanasak1412/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
