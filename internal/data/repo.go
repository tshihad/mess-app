package data

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// TODO add cache
type repo struct {
	logrus.FieldLogger
	*sql.DB
	// Cache
}

// NewRepo returns repo instance
func NewRepo(logger logrus.FieldLogger, db *sql.DB) *repo {
	return &repo{
		logger,
		db,
	}
}
