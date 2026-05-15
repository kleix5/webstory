package main

import (
	"fmt"
	"log"
	"net/http"
)

const Port = "8080"

func main() {
	// Статические файлы (CSS, JS, картинки) из папки static/
	// Браузер запросит /assets/css/main.css → сервер отдаст static/assets/css/main.css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/images/", http.FileServer(http.Dir("static")))

	InitDB()

	// Страницы
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/order", orderHandler)
	fmt.Println("Сервер запущен: http://localhost:" + Port)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
