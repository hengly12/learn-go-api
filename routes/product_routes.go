package routes

import (
	"go-api/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterProductRoutes(r chi.Router) {
	r.Route("/product", func(r chi.Router) {
		r.Get("/", handlers.GetProducts)
		r.Post("/", handlers.CreateProduct)
		r.Delete("/{id}", handlers.DeleteProduct)
	})
}
