package handler

import (
	"encoding/json"
	"internal/internal/dto"
	"internal/internal/repository"
	"log"
	"net/http"
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
