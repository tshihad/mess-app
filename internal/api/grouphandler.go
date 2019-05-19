package api

import "net/http"

func (a *App) handleAddGroup(w http.ResponseWriter, r *http.Request) {
	err := a.InsertGroup()
	if err != nil {
		a.Fail(w, 0, "Failed to create group", http.StatusConflict)
		return
	}
}
