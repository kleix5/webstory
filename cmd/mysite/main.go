package main

import (
	"fmt"
	"log"
	"net/http"
)

const Host = "http://localhost:5000"

func main() {
	// 1. Указываем, что отдавать файлы из текущей директории
	fs := http.FileServer(http.Dir("."))
	http.Handle("./", fs)

	port := "5000"
	url := Host

	fmt.Println("Сервер запущен:", url)

	http.Handle("/", http.FileServer(http.Dir(".")))
	// http.ListenAndServe(":8080", nil)

	// 3. Запуск сервера
	log.Fatal(http.ListenAndServe(":"+port, nil))

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		http.Redirect(w, r, "/layout.html", http.StatusSeeOther) // 303
	})
	// http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	tmpl, err := template.ParseFiles("layout.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	// 	tmpl.Execute(w, nil)
	// 	if err := tmpl.Execute(w, nil); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	// fmt.Fprintln(w, "POST received")
	// })
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
