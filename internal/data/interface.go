package data

import (
	"context"
	"mess-app/internal/models"
	"net/http"
	"time"
)

// Repo - Data repository
type Repo interface {
	UserRepo
	APIRespose
	TokenRepo
}

// UserRepo for user opertations
type UserRepo interface {
	InsertUser(ctx context.Context, user models.UserPayload) error
	GetUsers(ctx context.Context) ([]models.UserResponsePayload, error)
	GetUserByUserName(username string) (*models.UserPayload, error)
}

// Cache := TODO
type Cache interface {
	Get(key string, val interface{}) error
	Set(key string, val interface{}, d time.Duration) error
	IfExists(key string) (bool, error)
}

// APIRespose defined Success and Fail responses
type APIRespose interface {
	Success(w http.ResponseWriter, status int, data interface{})
	Fail(w http.ResponseWriter, errorCode int, message string, status int)
}
type TokenRepo interface {
	GenerateAuthToken(username string) (string, error)
	GetAuthToken(username string) (string, error)
	ValidateAuthToken(username, token string) error
	SetAuthToken(username, token string) error
}
