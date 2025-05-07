package models

import (
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type Question struct {
	ID     uuid.UUID `json:"id"`
	QuizID uuid.UUID `json:"quiz_id"`
	Text   string    `json:"text"`
}

type Answer struct {
	ID         uuid.UUID `json:"id"`
	QuestionID uuid.UUID `json:"question_id"`
	Text       string    `json:"text"`
	IsCorrect  bool      `json:"is_correct"`
}
