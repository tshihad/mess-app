package data

import (
	"context"
	"mess-app/internal/models"
	"net/http"
)

// Repo - Data repository
type Repo interface {
	UserRepo
	APIRespose
}

// UserRepo for user opertations
type UserRepo interface {
	InsertUser(ctx context.Context, user models.UserPayload) error
	GetUsers(ctx context.Context) (*models.UserPayload, error)
}

// Cache := TODO
type Cache interface {
}

// APIRespose defined Success and Fail responses
type APIRespose interface {
	Success(w http.ResponseWriter, status int, data interface{})
	Fail(w http.ResponseWriter, errorCode int, message string, status int)
}
