package models

import "time"

type Result struct {
	ID          int       `json:"id"`
	UserEmail   string    `json:"user_email"`
	QuizID      string    `json:"quiz_id"`
	Score       int       `json:"score"`
	Total       int       `json:"total"`
	SubmittedAt time.Time `json:"submitted_at"`
}
