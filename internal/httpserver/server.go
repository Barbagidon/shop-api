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

	mux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server := &Server{
		httpServer: httpServer,
		mux:        mux,
	}

	server.registerRoutes()

	return server
}

func (s *Server) Start() error {
	fmt.Printf("HTTP server started on %s\n", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("GET /health", s.health)
}
