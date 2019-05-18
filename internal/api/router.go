package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router for app
func (a *App) Router() http.Handler {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Post("/user", a.handlePostUser)
		r.Get("/user", a.handleGetUsers)
	})
	return r
}
