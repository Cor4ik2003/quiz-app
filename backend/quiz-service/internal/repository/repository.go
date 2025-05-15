package repository

import (
	"context"
	"internal/internal/db"
	"internal/internal/dto"
)

type Quiz struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

func CreateQuizWithQuestions(title, description string, questions []dto.Question) (string, error) {
	tx, err := db.DB.Begin(context.Background())
	if err != nil {
		return "", err
	}
	defer tx.Rollback(context.Background())

	var quizID string
	err = tx.QueryRow(context.Background(),
		`INSERT INTO quizzes (title, description) VALUES ($1, $2) RETURNING id`,
		title, description).Scan(&quizID)
	if err != nil {
		return "", err
	}

	for _, q := range questions {
		var questionID string
		err = tx.QueryRow(context.Background(),
			`INSERT INTO questions (quiz_id, text) VALUES ($1, $2) RETURNING id`,
			quizID, q.Text).Scan(&questionID)
		if err != nil {
			return "", err
		}

		for _, a := range q.Answers {
			_, err = tx.Exec(context.Background(),
				`INSERT INTO answers (question_id, text, is_correct) VALUES ($1, $2, $3)`,
				questionID, a.Text, a.IsCorrect)
			if err != nil {
				return "", err
			}
		}
	}

	err = tx.Commit(context.Background())
	return quizID, err
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

func GetQuizByID(ctx context.Context, id string) (Quiz, error) {
	var q Quiz
	err := db.DB.QueryRow(ctx, `SELECT id, title, created_by FROM quizzes WHERE id = $1`, id).
		Scan(&q.ID, &q.Title, &q.CreatedBy)
	return q, err
}

func DeleteQuiz(ctx context.Context, id string) error {
	_, err := db.DB.Exec(ctx, `DELETE FROM quizzes WHERE id = $1`, id)
	return err
}
