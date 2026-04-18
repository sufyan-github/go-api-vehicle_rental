package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to API",
		})
	})

	// Users
	r.GET("/users", func(c *gin.Context) {
		users := []string{"Sufyan", "John", "Alice"}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	r.POST("/user", func(c *gin.Context) {
	var body map[string]string

	c.BindJSON(&body)

	c.JSON(200, gin.H{
		"name": body["name"],
		"message": "User created",
	})
})

	r.Run(":8080")
}
