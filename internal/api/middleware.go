package api

import (
	"net/http"
	"strings"
)

func (a *App) middlewareAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		username := r.Header.Get("username")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			a.Fail(w, 0, "Invalied Authorization token format", http.StatusBadRequest)
			return
		}
		if err := a.ValidateAuthToken(username, splitToken[1]); err != nil {
			a.Fail(w, 0, "Invalid Auth token", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
