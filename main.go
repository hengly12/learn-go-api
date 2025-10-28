package main

import (
	"log"
	"net/http"

	"go-api/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
