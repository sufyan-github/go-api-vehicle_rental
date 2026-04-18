package main

import (
	"go-api/config"
	"go-api/middleware"
	"go-api/models"
	"go-api/routes"
	"go-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.ConnectRedis()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Booking{},
	)

	routes.SetupRoutes(r)
	go utils.StartConsumer()

	r.Use(middleware.RateLimiter())

	r.Run(":8080")
}