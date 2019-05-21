package data

import (
	"mess-app/internal/core"

	"github.com/pkg/errors"
)

func (r *repo) GenerateAuthToken(username string) (string, error) {
	return "", nil
}
func (r *repo) GetAuthToken(username string) (string, error) {
	var tok string
	err := r.Cache.Get(username, &tok)
	if err != nil {
		return "", errors.Wrap(err, "Failed to get token")
	}
	return tok, nil
}

// simple auth, TODO implement jwt
func (r *repo) ValidateAuthToken(username, token string) error {
	tok, err := r.GetAuthToken(username)
	if err != nil {
		return err
	}
	if tok != token {
		return core.ErrInvalidToken
	}
	return nil
}

func (r *repo) SetAuthToken(username, token string) error {
	return r.Cache.Set(username, token, 0)
}
