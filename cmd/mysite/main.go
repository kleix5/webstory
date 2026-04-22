package main

import (
	"fmt"
	"log"
	"net/http"
)

const Host = "http://localhost:8080"

func main() {
	// 1. Указываем, что отдавать файлы из текущей директории
	fs := http.FileServer(http.Dir("../../."))
	http.Handle("../../", fs)

	port := "8080"
	url := Host

	fmt.Println("Сервер запущен:", url)

	http.Handle("/", http.FileServer(http.Dir("../../.")))
	http.ListenAndServe(":8080", nil)

	// 3. Запуск сервера
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
//     // Парсим и выводим HTML-файл
//     tmpl, err := template.ParseFiles("index.html")
//     if err != nil {
//         http.Error(w, err.Error(), 400)
//         return
//     }
//     tmpl.Execute(w, nil)
// }

// Обработчик, который принимает POST-запрос от кнопки
func buttonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Обработка нажатия
		fmt.Println("Кнопка нажата пользователем!")

		// Отправляем ответ обратно клиенту
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Нажатие обработано"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
