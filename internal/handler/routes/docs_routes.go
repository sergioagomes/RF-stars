package routes

import (
	//_ "goauth/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	docsURL = "http://localhost:8080/docs/doc.json"
)

// @title		API FitTracker
// @version	1.0
func InitDocsRoutes(r chi.Router) {
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

}
