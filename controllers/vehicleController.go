package controllers

import (
	"go-api/models"
	"go-api/services"

	"github.com/gin-gonic/gin"
)

func CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	c.BindJSON(&vehicle)

	err := services.CreateVehicle(vehicle)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, vehicle)
}

func GetVehicles(c *gin.Context) {
	vehicles, _ := services.GetVehicles()
	c.JSON(200, vehicles)
}