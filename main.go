package main

import (
	"fmt"
	"log"
	"net/http"
)

const Port = "8080"

func main() {
	// Статические файлы  из папки templates/
	// Браузер запросит /templates/ → сервер отдаст static/css/style.css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	InitDB()

	// Страницы
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/admin/status", statusHandler)
	fmt.Println("Сервер запущен: http://localhost:" + Port)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
