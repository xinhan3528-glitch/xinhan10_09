package routes

import (
	"xinhan10_09/Internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/posts/{id}", app.PostHandler.HandlerGetPostById)
	r.Post("/posts", app.PostHandler.HandlerCreatePost)
	r.Get("/posts", app.PostHandler.HandleGetPosts)

	return r
}
