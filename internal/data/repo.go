package data

import (
	"database/sql"
	"mess-app/shared"

	"github.com/sirupsen/logrus"
)

// TODO add cache
type repo struct {
	logrus.FieldLogger
	*sql.DB
	shared.Response
	// Cache
}

// NewRepo returns repo instance
func NewRepo(logger logrus.FieldLogger, db *sql.DB, response shared.Response) *repo {
	return &repo{
		logger,
		db,
		response,
	}
}
