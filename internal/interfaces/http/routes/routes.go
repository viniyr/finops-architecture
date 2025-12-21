package routes

import (
	"aevyn/finops-arch/internal/interfaces/http/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/health", handlers.Health)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/chargebacks", handlers.OpenChargeback)
	})

	return r
}
