package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser хеширует пароль и создает пользователя
func RegisterUser(email, password, role string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("не удалось найти пользователя: %w", err)
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("неверный пароль")
	}

	// Генерация токена
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", fmt.Errorf("ошибка генерации токена: %w", err)
	}

	return token, nil
}
