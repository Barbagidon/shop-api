package httpserver

import "fmt"

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() {
	fmt.Print("HTTP server started")
}
