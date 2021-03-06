package api

import (
	"context"
	"encoding/json"
	"mess-app/internal/core"
	"mess-app/internal/models"
	"net/http"
)

// PostUser create new user
func (a *App) handlePostUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var user models.UserPayload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		a.Fail(w, core.JSON_ERROR_CODE, err.Error(), 500)
		return
	}
	if err := a.InsertUser(ctx, user); err != nil {
		a.Fail(w, 0, err.Error(), 500)
		return
	}
	// Removing password from response
	*user.Password = ""
	a.Success(w, 201, user)
}
