package repository

import (
	"context"
	"internal/internal/db"
)

type Quiz struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateQuiz(title, description string) (string, error) {
	query := `INSERT INTO public.quizzes (title, description) VALUES ($1, $2) RETURNING id`
	var id string
	err := db.DB.QueryRow(context.Background(), query, title, description).Scan(&id)
	return id, err
}

func GetAllQuizzes() ([]Quiz, error) {
	rows, err := db.DB.Query(context.Background(), `SELECT id, title, description FROM quizzes`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []Quiz
	for rows.Next() {
		var q Quiz
		if err := rows.Scan(&q.ID, &q.Title, &q.Description); err != nil {
			return nil, err
		}
		quizzes = append(quizzes, q)
	}
	return quizzes, nil
}
