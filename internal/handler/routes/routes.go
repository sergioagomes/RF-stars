package routes

import (
	"goauth/internal/handler"

	_ "goauth/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	docsURL = "http://localhost:8080/docs/doc.json"
)

// @title		Swagger Dark Mode
// @version	1.0
func InitRoutes(r chi.Router) {
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

	r.Get("/user", handler.GetUser)
}
