package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (a *App) Router() http.Handler {
	r := chi.NewRouter()
	return r
}
