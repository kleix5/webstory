package app

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	// Handle the Get Started form submit
	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.Redirect(w, r, "/layout.html", http.StatusSeeOther) // 303
		fmt.Println("The BUTTON was Clicked!!!")
	})

	return &Server{mux: mux}
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) routes() {
	s.mux.Handle("/", http.FileServer(http.Dir(".")))
}
