package main

import (
	"html/template"
	"log"
	"net/http"
)

var counterClick int64

//Главная страница

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//Любой путь кроме "/" отправляем на 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}

// func clickHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	count := atomic.AddInt64(&counterClick, 1)
// 	fmt.Printf("Кнопка нажата: %d раз\n", count)

// 	http.Redirect(w, r, "/thanks", http.StatusSeeOther)
// }

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	contact := r.FormValue("contact")
	message := r.FormValue("message")

	_, err := DB.Exec(
		"INSERT INTO orders (name, contact, message) VALUES ($1, $2, $3)",
		name, contact, message,
	)

	if err != nil {
		log.Println("Ошибка", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
