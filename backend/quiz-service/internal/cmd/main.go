package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"internal/internal/db"
	"internal/internal/handler"
	"internal/internal/middleware"
)

func main() {

	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/quizzes", handler.GetQuizzes).Methods(http.MethodGet)
	r.HandleFunc("/quizzes/{id}", handler.GetQuizByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/quizzes/{id}/full", handler.GetFullQuiz).Methods(http.MethodGet)

	r.Handle("/quizzes", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateQuizHandler))).Methods(http.MethodPost)

	r.Handle("/quizzes/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.UpdateQuizHandler))).Methods(http.MethodPut)

	r.Handle("/quizzes/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.DeleteQuizHandler))).Methods(http.MethodDelete)

	r.Handle("/quizzes/{id}/attempt", middleware.AuthMiddleware(http.HandlerFunc(handler.SubmitAttemptHandler))).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Или явно указать: http://localhost:5500
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	//r.Handle("/attempts", middleware.AuthMiddleware(http.HandlerFunc(handler.GetUserAttemptsHandler))).Methods(http.MethodGet)

	handlerWithCors := c.Handler(r)

	log.Println("Quiz service started on :8082")
	log.Fatal(http.ListenAndServe(":8082", handlerWithCors))
}
