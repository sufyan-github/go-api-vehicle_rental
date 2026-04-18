package routes

import (
	"go-api/controllers"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// Auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Public
	r.GET("/vehicles", controllers.GetVehicles)

	// Protected
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/profile", controllers.Profile)
	auth.POST("/vehicles", controllers.CreateVehicle)
	auth.POST("/bookings", controllers.CreateBooking)
	auth.GET("/bookings", controllers.GetBookings)
	auth.PUT("/bookings/:id/cancel", controllers.CancelBooking)
}
