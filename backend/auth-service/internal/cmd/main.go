package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Auth Service is running"})
	})

	port := "8080"
	log.Printf("Auth Service is running on port %s", port)

	//r.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"status": "Auth Service is running"})
	//})

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
