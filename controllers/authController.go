package controllers

import (
	"go-api/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input map[string]string

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := services.RegisterUser(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input map[string]string

	c.BindJSON(&input)

	token, err := services.LoginUser(input)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}