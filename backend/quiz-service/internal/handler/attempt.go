package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"internal/internal/db"

	"github.com/gorilla/mux"
)

type AnswerSubmission struct {
	QuestionID string `json:"question_id"`
	AnswerID   string `json:"answer_id"`
}

type AttemptRequest struct {
	Answers []AnswerSubmission `json:"answers"`
}

func SubmitAttemptHandler(w http.ResponseWriter, r *http.Request) {
	quizID := mux.Vars(r)["id"]
	userID := r.Context().Value("user_id").(string)

	var req AttemptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}


	correctCount := 0
	for _, a := range req.Answers {
		var isCorrect bool
		query := `SELECT is_correct FROM answers WHERE id = $1 AND question_id = $2 AND is_correct = true`
		err := db.DB.QueryRow(context.Background(), query, a.AnswerID, a.QuestionID).Scan(&isCorrect)
		if err == nil && isCorrect {
			correctCount++
		}
	}

	// Сохраняем попытку
	_, err := db.DB.Exec(context.Background(),
		`INSERT INTO attempts (quiz_id, user_id, score) VALUES ($1, $2, $3)`,
		quizID, userID, correctCount,
	)
	if err != nil {
		http.Error(w, "Failed to save attempt", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"score": %d, "total": %d}`, correctCount, len(req.Answers))
}
