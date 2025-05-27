package repository

import (
	"context"
	"time"

	"internal/internal/db"
)

func SaveResult(email, quizID string, score, total int) error {
	_, err := db.DB.Exec(
		context.Background(),
		`INSERT INTO results (user_email, quiz_id, score, total, submitted_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		email, quizID, score, total, time.Now(),
	)
	return err
}
