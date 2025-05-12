package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"internal/internal/db"

	"github.com/gorilla/mux"
)

type FullAnswer struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type FullQuestion struct {
	ID      string       `json:"id"`
	Text    string       `json:"text"`
	Answers []FullAnswer `json:"answers"`
}

type FullQuiz struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Questions   []FullQuestion `json:"questions"`
}

func GetFullQuiz(w http.ResponseWriter, r *http.Request) {
	quizID := mux.Vars(r)["id"]

	// Получаем квиз
	var quiz FullQuiz
	err := db.DB.QueryRow(context.Background(),
		`SELECT id, title, description FROM quizzes WHERE id = $1`, quizID).
		Scan(&quiz.ID, &quiz.Title, &quiz.Description)
	if err != nil {
		http.Error(w, "Quiz not found", http.StatusNotFound)
		return
	}

	// Получаем вопросы
	qrows, err := db.DB.Query(context.Background(),
		`SELECT id, text FROM questions WHERE quiz_id = $1`, quizID)
	if err != nil {
		http.Error(w, "Failed to load questions", http.StatusInternalServerError)
		return
	}
	defer qrows.Close()

	for qrows.Next() {
		var q FullQuestion
		if err := qrows.Scan(&q.ID, &q.Text); err != nil {
			continue
		}

		// Получаем ответы для каждого вопроса
		arows, err := db.DB.Query(context.Background(),
			`SELECT id, text, is_correct FROM answers WHERE question_id = $1`, q.ID)
		if err != nil {
			continue
		}

		for arows.Next() {
			var a FullAnswer
			if err := arows.Scan(&a.ID, &a.Text, &a.IsCorrect); err == nil {
				q.Answers = append(q.Answers, a)
			}
		}
		arows.Close()

		quiz.Questions = append(quiz.Questions, q)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}
