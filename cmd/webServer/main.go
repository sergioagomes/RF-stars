package main

import (
	"fmt"
	"goauth/internal/handler/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	routes.InitRoutes(r)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
