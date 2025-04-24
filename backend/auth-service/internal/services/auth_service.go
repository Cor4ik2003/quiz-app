package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser хеширует пароль и создает пользователя
func RegisterUser(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	user := &models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	err = repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
