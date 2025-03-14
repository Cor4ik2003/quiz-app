package repository

import (
	"gorm.io/gorm"
)

// User – модель пользователя
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// UserRepository – интерфейс репозитория пользователей
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}

// userRepo – реализация репозитория
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository – конструктор репозитория
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// Create – добавление нового пользователя
func (r *userRepo) Create(user *User) error {
	return r.db.Create(user).Error
}

// FindByEmail – поиск пользователя по email
func (r *userRepo) FindByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
