package main

import (
	"fmt"
	"log"
	"net/http"
)

const Host = "http://localhost:8080"

var counter_click int

func main() {
	// Указываем, что отдавать файлы из текущей директории
	http.Handle("/", http.FileServer(http.Dir("./")))
	port := "8080"
	url := Host

	fmt.Println("Сервер запущен:", url)

	http.HandleFunc("/click", clickHandler)
	// Запуск сервера
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Обработчик, который принимает POST-запрос от кнопки
func clickHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	counter_click++
	fmt.Println("Кнопка нажата:", counter_click, "раз")

	http.Redirect(w, r, "layout.html", http.StatusSeeOther)
}
