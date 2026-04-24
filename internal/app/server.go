package app

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
