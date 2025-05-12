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

func GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, role, created_at
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при поиске пользователя по email: %w", err)
	}

	return &user, nil
}
