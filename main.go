package main

import (
	"go-api/config"
	"go-api/models"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Booking{},
	)

	routes.SetupRoutes(r)

	r.Run(":8080")
}