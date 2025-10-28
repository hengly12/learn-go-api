package routes

import (
    "github.com/go-chi/chi/v5"
    "go-api/handlers"
)

func RegisterUserRoutes(r chi.Router) {
    r.Route("/users", func(r chi.Router) {
        r.Get("/", handlers.GetUsers)
        r.Get("/{id}", handlers.GetUser)
        r.Post("/", handlers.CreateUser)
        r.Put("/{id}", handlers.UpdateUser)
        r.Delete("/{id}", handlers.DeleteUser)
        r.Post("/login", handlers.LoginUser)
    })
}
