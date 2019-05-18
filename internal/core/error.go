package core

import "errors"

const (
	// JSON_ERROR_CODE - json unmarshalling and marshalling
	// error code for api response
	JSON_ERROR_CODE = 1001
)

// ErrInvalidPassword invalid password error
var ErrInvalidPassword = errors.New("Invalid password")
