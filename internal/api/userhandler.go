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
		if err == core.ErrConflict {
			a.Fail(w, core.CONFLICT_ERROR_CODE, "User already present", http.StatusFound)
			return
		}
		a.Fail(w, 0, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := models.UserResponsePayload{
		UserName: user.UserName,
		Email:    user.Email,
		UserType: user.UserType,
	}
	a.Success(w, http.StatusCreated, resp)
}

// Deprecated function
func (a *App) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	users, err := a.GetUsers(ctx)
	if err != nil {
		a.Fail(w, 0, err.Error(), http.StatusInternalServerError)
		return
	}
	a.Success(w, http.StatusOK, users)
}
