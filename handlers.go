package main

import (
	"html/template"
	"log"
	"net/http"
)

type Order struct {
	Id        int
	Name      string
	Contact   string
	Message   string
	Status    string
	CreatedAt string
}

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

func adminHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		return
	}

	var orders []Order
	rows, err := DB.Query("SELECT id, name, contact, message, status, created_at FROM orders")

	for rows.Next() {
		var o Order
		rows.Scan(&o.Id, &o.Name, &o.Contact, &o.Message, &o.Status, &o.CreatedAt)
		orders = append(orders, o)
	}

	tmpl.Execute(w, orders)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	status := r.FormValue("status")
	DB.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, id)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
