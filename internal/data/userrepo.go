package data

import (
	"context"
	"mess-app/internal/core"
	"mess-app/internal/models"

	"github.com/pkg/errors"
)

var userInsertQuery = `INSERT INTO users ( users_name, email, password_digest, users_type) VALUES($1,$2,$3,$4)`

func (r *repo) InsertUser(ctx context.Context, user models.UserPayload) error {
	var userType int
	switch *user.UserType {
	case core.ADMIN_KEY:
		userType = core.ADMIN
	case core.COLLEGUES_KEY:
		userType = core.COLLEGUES
	case core.COOK_KEY:
		userType = core.COOK
	}
	hashedPwd, err := core.HashPassword([]byte(*user.Password))
	if err != nil {
		return errors.Wrap(err, "Failed to hash password")
	}
	res, err := r.DB.Exec(userInsertQuery, *user.Email, *user.UserName, hashedPwd, userType)
	if err != nil {
		return errors.Wrap(err, "Failed to insert user")
	}
	if l, err := res.RowsAffected(); err != nil || l != 1 {
		return errors.New("Failed to insert user")
	}
	return nil
}
func (r *repo) GetUsers(ctx context.Context) (*models.UserPayload, error) {
	return nil, nil
}
