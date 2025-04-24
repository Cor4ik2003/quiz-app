package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Хранится в зашифрованном виде
	CreatedAt time.Time `json:"created_at"`
}
