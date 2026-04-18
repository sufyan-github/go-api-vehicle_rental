package main

import (
	"log"
	"net/http"
	"strconv"

	"go-api/config"
	"go-api/middleware"
	"go-api/models"
	"go-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Booking{},
	); err != nil {
		log.Fatal("Migration failed:", err)
	}

	setupRoutes(r)

	r.Run(":8080")
}

func setupRoutes(r *gin.Engine) {

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	userRoutes := r.Group("/users")

	// Register
	userRoutes.POST("/register", func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hash)

		config.DB.Create(&user)

		c.JSON(201, gin.H{"message": "User registered"})
	})

	// Login
	userRoutes.POST("/login", func(c *gin.Context) {
		var input models.User
		var user models.User

		c.BindJSON(&input)

		config.DB.Where("email = ?", input.Email).First(&user)

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		token, _ := utils.GenerateToken(user.ID)

		c.JSON(200, gin.H{"token": token})
	})

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	// Profile
	protected.GET("/profile", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	// Create Vehicle
	protected.POST("/vehicles", func(c *gin.Context) {
		var v models.Vehicle
		c.BindJSON(&v)

		v.Available = true
		config.DB.Create(&v)

		c.JSON(201, v)
	})

	// Get Vehicles
	r.GET("/vehicles", func(c *gin.Context) {
		var v []models.Vehicle
		config.DB.Find(&v)
		c.JSON(200, v)
	})

	// Create Booking
	protected.POST("/bookings", func(c *gin.Context) {
		var booking models.Booking
		c.BindJSON(&booking)

		userID, _ := c.Get("user_id")
		booking.UserID = uint(userID.(float64))

		var vehicle models.Vehicle
		config.DB.First(&vehicle, booking.VehicleID)

		if !vehicle.Available {
			c.JSON(400, gin.H{"error": "Vehicle not available"})
			return
		}

		vehicle.Available = false
		config.DB.Save(&vehicle)

		booking.Status = "booked"
		config.DB.Create(&booking)

		c.JSON(201, booking)
	})

	// Get My Bookings
	protected.GET("/bookings", func(c *gin.Context) {
		userID, _ := c.Get("user_id")

		var bookings []models.Booking
		config.DB.Where("user_id = ?", userID).Find(&bookings)

		c.JSON(200, bookings)
	})

	// Cancel Booking
	protected.PUT("/bookings/:id/cancel", func(c *gin.Context) {
		id := c.Param("id")

		var booking models.Booking
		config.DB.First(&booking, id)

		booking.Status = "cancelled"
		config.DB.Save(&booking)

		var vehicle models.Vehicle
		config.DB.First(&vehicle, booking.VehicleID)

		vehicle.Available = true
		config.DB.Save(&vehicle)

		c.JSON(200, gin.H{"message": "Booking cancelled"})
	})
}