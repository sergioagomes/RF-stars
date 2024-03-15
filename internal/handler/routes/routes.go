package routes

import (
	"goauth/internal/handler"

	"github.com/go-chi/chi/v5"
)

// @title		Swagger Dark Mode
// @version	1.0
func InitRoutes(r chi.Router) {
	r.Get("/user", handler.GetUser)
}
