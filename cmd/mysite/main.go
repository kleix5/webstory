package main

import (
	"fmt"
	"log"
	"net/http"
	// "log/slog"
	"flag"
	"os"

)

const Host = "http://localhost:5000"

func main() {
	addr := flag.String("addr", ":5000", "Сетевой адрес веб-сервера")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "СООБЩЕНИЕ ОБ ОТРАБОТКЕ")
	})
	// mux.HandleFunc("/snippet", showSnippet)
	// mux.HandleFunc("/snippet/create", createSnippet)
 
	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
 
	// Применяем созданные логгеры к нашему приложению.
	infoLog.Printf("Запуск сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)

	// // 1. Указываем, что отдавать файлы из текущей директории
	// fs := http.FileServer(http.Dir("."))
	// http.Handle("./", fs)

	// port := "5000"
	// // url := Host

	// // http.HandleFunc("./", handler)
	// slog.Info("Сервер запущен", "port", 5000)

	// http.Handle("/", http.FileServer(http.Dir(".")))
	// if err != nil {
  	// 	log.Printf("Ошибка: %v", err) // Запишет в терминал и продолжит работу
	// }

	// // 3. Запуск сервера
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}


// func handler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Запрос: %s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
// }
func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("Обработан запрос:", r.URL.Path) // Лог в терминал
    w.Write([]byte("Hello, world!"))
}


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


// func handler(w http.ResponseWriter, r *http.Request) {
//     // Парсим и выводим HTML-файл
//     tmpl, err := template.ParseFiles("index.html")
//     if err != nil {
//         http.Error(w, err.Error(), 400)
//         return
//     }
//     tmpl.Execute(w, nil)
// }
