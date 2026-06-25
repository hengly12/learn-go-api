package api

import (
	"net/http"

	"go-api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	r.ServeHTTP(w, req)
}
