package handler

import (
	"encoding/json"
	"internal/internal/repository"
	"log"
	"net/http"
)

type createQuizRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var req createQuizRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := repository.CreateQuiz(req.Title, req.Description)
	if err != nil {
		log.Println("Failed to create quiz:", err)
		http.Error(w, "Failed to create quiz", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzes, err := repository.GetAllQuizzes()
	if err != nil {
		log.Println("Error in GetAllQuizzes:", err)
		http.Error(w, "Failed to get quizzes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(quizzes)
}
