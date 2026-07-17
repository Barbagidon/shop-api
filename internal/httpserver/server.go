package httpserver

import (
	"fmt"
	"net/http"

	"github.com/Barbagidon/shop-api/internal/service"
)

type Server struct {
	httpServer     *http.Server
	mux            *http.ServeMux
	productService *service.ProductService
}

func New(s *service.ProductService) *Server {

	mux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server := &Server{
		httpServer:     httpServer,
		mux:            mux,
		productService: s,
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
