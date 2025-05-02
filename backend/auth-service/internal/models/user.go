package models

import (
	"time"

	"github.com/google/uuid" // Импорт для UUID
)

type User struct {
	ID        uuid.UUID `json:"id"` // изменяем на uuid.UUID
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Хранится в зашифрованном виде
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
