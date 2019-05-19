package data

import (
	"context"
	"mess-app/internal/core"
	"mess-app/internal/models"

	"github.com/pkg/errors"
)

var (
	userInsertQuery = `INSERT INTO users ( users_name, email, password_digest, users_type) VALUES($1,$2,$3,$4)`
	getusersQuery   = `SELECT users_name,email,users_type FROM users;`
)

func (r *repo) InsertUser(ctx context.Context, user models.UserPayload) error {
	var userType int
	switch *user.UserType {
	case core.ADMIN_KEY:
		userType = core.ADMIN
	case core.COLLEGUES_KEY:
		userType = core.COLLEGUES
	case core.COOK_KEY:
		userType = core.COOK
	default:
		return core.ErrInvalidUserType
	}
	hashedPwd, err := core.HashPassword([]byte(*user.Password))
	if err != nil {
		return errors.Wrap(err, "Failed to hash password")
	}
	res, err := r.DB.Exec(userInsertQuery, *user.UserName, *user.Email, hashedPwd, userType)
	if err != nil {
		return core.ErrConflict
	}
	if l, err := res.RowsAffected(); err != nil || l != 1 {
		return errors.New("Failed to insert user")
	}
	return nil
}

// GetUsers Deprecated
func (r *repo) GetUsers(ctx context.Context) ([]models.UserResponsePayload, error) {
	var users []models.UserResponsePayload
	var userType int
	var userTypeString, userName, email string
	rows, err := r.DB.Query(getusersQuery)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch users")
	}
	for rows.Next() {
		if err := rows.Scan(&userName, &email, &userType); err != nil {
			return nil, errors.Wrap(err, "Failed to scan rows")
		}
		switch userType {
		case core.ADMIN:
			userTypeString = core.ADMIN_KEY
		case core.COLLEGUES:
			userTypeString = core.COLLEGUES_KEY
		case core.COOK:
			userTypeString = core.COOK_KEY
		}
		users = append(users, models.UserResponsePayload{
			UserName: &userName,
			Email:    &email,
			UserType: &userTypeString,
		})
	}
	return users, nil
}

func (r *repo) GetUserByUserName(username string) (*models.UserPayload, error) {
	return nil, nil
}
