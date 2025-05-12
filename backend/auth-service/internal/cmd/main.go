package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"
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
	r.POST("/login", handlers.LoginHandler) // ✅ Добавляем логин

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// Только для студентов
	auth.GET("/tests", middleware.RoleMiddleware("student"), handlers.PassTestHandler)

	// Только для преподавателей
	auth.POST("/tests", middleware.RoleMiddleware("teacher"), handlers.CreateTestHandler)

	log.Println("Auth service started on port 8081")
	r.Run(":8081")
}
