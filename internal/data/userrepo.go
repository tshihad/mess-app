package data

import (
	"context"
	"mess-app/internal/core"
	"mess-app/internal/models"

	"github.com/pkg/errors"
)

var userInsertQuery = `INSERT INTO users ( users_name, email, users_type) VALUES($1,$2,$3)`

func (r *repo) InsertUser(ctx context.Context, user *models.UserPayload) error {
	var userType int
	switch *user.UserType {
	case core.ADMIN_KEY:
		userType = core.ADMIN
	case core.COLLEGUES_KEY:
		userType = core.COLLEGUES
	case core.COOK_KEY:
		userType = core.COOK
	}
	res, err := r.DB.Exec(userInsertQuery, *user.UserName, *user.Email, userType)
	if err != nil {
		return errors.Wrap(err, "Failed to insert user")
	}
	if l, err := res.LastInsertId(); err != nil || l != 1 {
		return errors.Wrap(err, "Failed to insert user")
	}
	return nil
}
