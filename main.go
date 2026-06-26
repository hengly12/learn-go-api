// package main

// import (
// 	"log"
// 	"net/http"

// 	"go-api/routes"

// 	"github.com/go-chi/chi/v5"
// )

// func main() {
// 	r := chi.NewRouter()

// 	routes.RegisterUserRoutes(r)
// 	routes.RegisterProductRoutes(r)

// 		log.Println("🚀 Server running on http://localhost:8080")
// 		http.ListenAndServe(":8080", r)
// 	}

// run code in terminal air or go run main.go

package main

import (
	"log"
	"net/http"
	"os"

	"go-api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // or []string{"*"} for all
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
