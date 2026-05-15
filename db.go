package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Строка подключения к PostgreSQL
	connStr := "user=postgres password=Go1111Go dbname=webstory sslmode=disable"
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
