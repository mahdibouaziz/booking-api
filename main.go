package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdibouaziz/booking-api/db"
	"github.com/mahdibouaziz/booking-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
