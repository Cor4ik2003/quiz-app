package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool 

func Init() *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	DB = pool 
	log.Println("Database connection initialized")
	return pool
}
