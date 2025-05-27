package repository

import (
	"context"
	"fmt"
	"internal/internal/db"
	model "internal/internal/model"
)

func CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, role, created_at)
		VALUES (gen_random_uuid(), $1, $2, $3, NOW())
		RETURNING id, created_at
	`
	return db.DB.QueryRow(context.Background(), query, user.Email, user.Password, user.Role).
		Scan(&user.ID, &user.CreatedAt)
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, email, password_hash, role, created_at FROM users WHERE email = $1`

	err := db.DB.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %w", err)
	}
	return &user, nil
}
