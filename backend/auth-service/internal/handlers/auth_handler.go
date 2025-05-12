package handlers

import (
	service "auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=student teacher"`
}

func PassTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Прохождение теста доступно студенту"})
}

func CreateTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Создание теста доступно преподавателю"})
}

func RegisterHandler(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный ввод: " + err.Error()})
		return
	}

	user, err := service.RegisterUser(input.Email, input.Password, input.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный ввод: " + err.Error()})
		return
	}

	token, err := service.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
