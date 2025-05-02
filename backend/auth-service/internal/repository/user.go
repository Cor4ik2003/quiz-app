package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"fmt"
	// Импорт для работы с UUID
)

// CreateUser создает нового пользователя в базе данных
func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, role, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, created_at
	`

	err := config.DB.QueryRow(query, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %w", err)
	}

	return nil
}
