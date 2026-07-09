package httpserver

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

func New() *Server {

	httpServer := &http.Server{
		Addr: ":8080",
	}

	mux := http.NewServeMux()

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	fmt.Println("HTTP server started on :8080")

	return s.httpServer.ListenAndServe()
}
