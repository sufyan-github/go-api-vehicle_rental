package controllers

import (
	"go-api/models"
	"go-api/services"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	c.BindJSON(&booking)

	userID, _ := c.Get("user_id")

	result, err := services.CreateBooking(uint(userID.(float64)), booking)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}

func GetBookings(c *gin.Context) {
	userID, _ := c.Get("user_id")

	bookings, _ := services.GetBookings(userID)
	c.JSON(200, bookings)
}

func CancelBooking(c *gin.Context) {
	id := c.Param("id")

	err := services.CancelBooking(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Booking cancelled"})
}