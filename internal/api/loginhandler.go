package api

import (
	"mess-app/internal/core"
	"net/http"
)

func (a *App) handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := a.GetUserByUserName(username)
	if err != nil {
		a.Fail(w, 0, "Invalid username or password", http.StatusUnauthorized)
		a.Error("Invalid username or password ", err)
		return
	}
	err = core.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		a.Fail(w, 0, "Invalid username or password", http.StatusUnauthorized)
		a.Error("Invalid username or password ", err)
		return
	}
	token, err := a.GenerateAuthToken(username)
	if err != nil {
		a.Fail(w, 0, "Something went wrong, Can't create Token", http.StatusInternalServerError)
		a.Error(err, " Token error")
		return
	}
	// TODO
	a.SetAuthToken(username, token)
	a.Success(w, http.StatusCreated, token)
}
