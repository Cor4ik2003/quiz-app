package handler

import (
	"context"
	"encoding/json"
	"internal/internal/db"
	"internal/internal/dto"
	"internal/internal/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzes, err := repository.GetAllQuizzes()
	if err != nil {
		log.Println("Error in GetAllQuizzes:", err)
		http.Error(w, "Failed to get quizzes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(quizzes)
}

func CreateQuizHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQuizRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := repository.CreateQuizWithQuestions(req.Title, req.Description, req.Questions)
	if err != nil {
		log.Println("Failed to create quiz:", err)
		http.Error(w, "Failed to create quiz", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

type Question struct {
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

func DeleteQuizHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(string)
	quizID := mux.Vars(r)["id"]

	// Проверка владельца
	var createdBy string
	err := db.DB.QueryRow(context.Background(), `SELECT created_by FROM quizzes WHERE id = $1`, quizID).Scan(&createdBy)
	if err != nil {
		http.Error(w, "Quiz not found", http.StatusNotFound)
		return
	}
	if createdBy != userID {
		http.Error(w, "Forbidden: not your quiz", http.StatusForbidden)
		return
	}

	_, err = db.DB.Exec(context.Background(), `DELETE FROM quizzes WHERE id = $1`, quizID)
	if err != nil {
		http.Error(w, "Failed to delete quiz", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetQuizByIDHandler(w http.ResponseWriter, r *http.Request) {
	quizID := mux.Vars(r)["id"]

	quiz, err := repository.GetQuizByID(r.Context(), quizID)
	if err != nil {
		log.Println("Error in GetQuizByID:", err)
		http.Error(w, "Quiz not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}
