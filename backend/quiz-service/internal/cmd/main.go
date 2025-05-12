package main

import (
	"internal/internal/db"
	"internal/internal/handler"
	"log"
	"net/http"
)

func main() {
	// Инициализация подключения к БД
	db.Init()

	http.HandleFunc("/quizzes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetQuizzes(w, r)
		case http.MethodPost:
			handler.CreateQuizHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Quiz service started on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
