package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	port := "3000"

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port)
}
