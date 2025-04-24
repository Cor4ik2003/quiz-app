package repository

import (
	"auth-service/internal/models"
	"database/sql"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Создание пользователя
func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (email, password, created_at) VALUES ($1, $2, NOW())",
		user.Email, user.Password)
	return err
}

// Получение пользователя по email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.DB.QueryRow("SELECT id, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("пользователь не найден")
	}

	return user, err
}
