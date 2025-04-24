package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("БД не отвечает:", err)
	}

	DB = db
	log.Println("Успешное подключение к базе данных")

	fmt.Println("host =", host)
	fmt.Println("port =", port)
	fmt.Println("user =", user)
	fmt.Println("password =", password)
	fmt.Println("dbname =", dbname)

}
