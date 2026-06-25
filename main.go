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

package main

import (
	"log"
	"net/http"
	"os"

	"go-api/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
