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
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", a.handleLogin)

		r.Route("/{username}", func(r chi.Router) {
			// TODO
			// r.Use(a.middlewareAuth)
			r.Post("/user", a.handlePostUser)
			r.Get("/user", a.handleGetUsers)
		})
	})
	return r
}
