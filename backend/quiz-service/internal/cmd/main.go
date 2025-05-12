package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"internal/internal/db"
	"internal/internal/handler"
	"internal/internal/middleware"
)

func main() {

	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/quizzes/{id}/full", handler.GetFullQuiz).Methods(http.MethodGet)

	r.HandleFunc("/quizzes", handler.GetQuizzes).Methods(http.MethodGet)

	r.Handle("/quizzes", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateQuizHandler))).Methods(http.MethodPost)

	r.Handle("/quizzes/{id}/attempt", middleware.AuthMiddleware(http.HandlerFunc(handler.SubmitAttemptHandler))).Methods(http.MethodPost)

	log.Println("Quiz service started on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
