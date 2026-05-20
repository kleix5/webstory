package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "Go1111Go")
	dbname := getEnv("DB_NAME", "webstory")

	var err error
	// Строка подключения к PostgreSQL
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, user, password, dbname,
	)
	//Открытие соединения с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Неверная конфигурация базы данных")
	}
	if err = db.Ping(); err != nil {
		log.Fatal("БД недоступна")
	}
	DB = db
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
