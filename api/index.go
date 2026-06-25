package api

import (
	"net/http"

	"go-api/routes"

	"github.com/go-chi/chi/v5"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	r := chi.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	r.ServeHTTP(w, req)
}
