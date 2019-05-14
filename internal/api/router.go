package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router for app
func (a *App) Router() http.Handler {
	r := chi.NewRouter()
	return r
}
