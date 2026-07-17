package app

import (
	"github.com/Barbagidon/shop-api/internal/httpserver"
	"github.com/Barbagidon/shop-api/internal/repository/memory"
	"github.com/Barbagidon/shop-api/internal/service"
)

type App struct {
	server *httpserver.Server
}

func New() *App {
	repo := memory.NewProductRepository()
	productService := service.NewProductService(repo)

	server := httpserver.New(productService)
	return &App{
		server: server,
	}
}

func (a *App) Run() error {
	return a.server.Start()
}
