package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router for app
func (a *App) Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	// TODO
	// r.Use(a.middlewareAuth)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/user", a.handlePostUser)
		r.Get("/user", a.handleGetUsers)
	})
	return r
}
