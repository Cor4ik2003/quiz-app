package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/register", handlers.RegisterHandler)

	log.Println("Auth service started on port 808")
	r.Run(":8081")
}
