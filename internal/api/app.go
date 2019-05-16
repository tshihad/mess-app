package api

import (
	"database/sql"
	"mess-app/internal/data"
	"mess-app/shared"

	"github.com/sirupsen/logrus"
)

// App serves api
type App struct {
	logrus.FieldLogger
	data.Repo
}

// NewApp return app instance
func NewApp(logger logrus.FieldLogger, db *sql.DB) *App {
	response := shared.NewResponse(logger)
	repo := data.NewRepo(logger, db, response)
	return &App{
		logger,
		repo,
	}
}
