package main

import (
	"log"
	"net/http"
	"strconv"

	"go-api/config"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	r := gin.Default()

	// Connect to Database
	config.ConnectDB()

	// Auto migrate models
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	// Setup routes
	setupRoutes(r)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}

func setupRoutes(r *gin.Engine) {

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	// Group: Users
	userRoutes := r.Group("/users")
	{
		// Create User
		userRoutes.POST("/", func(c *gin.Context) {
			var user models.User

			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			config.DB.Create(&user)

			c.JSON(http.StatusCreated, user)
		})

		// Get All Users
		userRoutes.GET("/", func(c *gin.Context) {
			var users []models.User

			config.DB.Find(&users)

			c.JSON(http.StatusOK, users)
		})

		// Get Single User
		userRoutes.GET("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user ID",
				})
				return
			}

			var user models.User

			if err := config.DB.First(&user, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}

			c.JSON(http.StatusOK, user)
		})

		// Update User
		userRoutes.PUT("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user ID",
				})
				return
			}

			var user models.User

			if err := config.DB.First(&user, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}

			var input models.User
			if err := c.BindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			user.Name = input.Name
			user.Email = input.Email

			config.DB.Save(&user)

			c.JSON(http.StatusOK, user)
		})

		// Delete User
		userRoutes.DELETE("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user ID",
				})
				return
			}

			if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to delete user",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
			})
		})
	}
}