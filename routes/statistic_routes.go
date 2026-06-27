package routes

import (
	"go-api/handlers"
	"go-api/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterStatisticRoutes(r chi.Router) {
	r.Route("/statistic", func(r chi.Router) {

		// Require token for all statistic endpoints
		r.Use(middleware.AuthMiddleware)
		r.Get("/products", handlers.GetStatistics)
	})
}
