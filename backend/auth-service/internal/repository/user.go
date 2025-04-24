package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"fmt"
)

// CreateUser создает нового пользователя в базе данных
func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, created_at)
		VALUES ($1, $2, NOW())
		RETURNING id, created_at
	`

	err := config.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %w", err)
	}

	return nil
}
