package app

import (
	"github.com/Barbagidon/shop-api/internal/httpserver"
)

type App struct {
	server *httpserver.Server
}

func New() *App {
	server := httpserver.New()
	return &App{
		server: server,
	}
}

func (a *App) Run() error {
	return a.server.Start()
}
