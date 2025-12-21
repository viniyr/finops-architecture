package application

import (
	"net/http"

	"aevyn/finops-arch/internal/interfaces/http/routes"
)

type App struct {
	Router http.Handler
}

func New() *App {
	router := routes.SetupRoutes()

	return &App{
		Router: router,
	}
}
