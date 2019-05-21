package core

import "errors"

const (
	// JSON_ERROR_CODE - json unmarshalling and marshalling
	// error code for api response
	JSON_ERROR_CODE     = 1001
	CONFLICT_ERROR_CODE = 4004
)

var (
	// ErrInvalidPassword invalid password error
	ErrInvalidPassword = errors.New("Invalid password")
	// ErrInvalidUserType invalid user type
	ErrInvalidUserType  = errors.New("Invalid user type.(admin, cook and collegues are allowed)")
	ErrConflict         = errors.New("Resource already exists")
	ErrResourceNotFound = errors.New("resource not found")
	ErrInvalidToken     = errors.New("Invalid token")
)
