package api

import (
	"database/sql"
	"mess-app/internal/data"

	"github.com/sirupsen/logrus"
)

// App serves api
type App struct {
	logrus.FieldLogger
	data.Repo
}

// NewApp return app instance
func NewApp(logger logrus.FieldLogger, db *sql.DB) *App {
	repo := data.NewRepo(logger, db)
	return &App{
		logger,
		repo,
	}
}
