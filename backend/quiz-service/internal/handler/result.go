package handler

import (
	"encoding/json"
	"internal/internal/repository"
	"internal/internal/utils"
	"net/http"
)

func SubmitResultHandler(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractUserFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		QuizID string `json:"quiz_id"`
		Score  int    `json:"score"`
		Total  int    `json:"total"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := repository.SaveResult(user.Email, input.QuizID, input.Score, input.Total); err != nil {
		http.Error(w, "Failed to save result", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
