package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/properties", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from API!",
		})
	})

	r.Run(":8080")
}