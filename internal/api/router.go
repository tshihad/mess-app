package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router for app
func (a *App) Router() http.Handler {
	r := chi.NewRouter()
	r.Post("/user", a.handlePostUser)
	return r
}
