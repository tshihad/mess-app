package api

import (
	"database/sql"
	"mess-app/internal/data"
	"mess-app/shared"

	"github.com/go-redis/redis"

	"github.com/sirupsen/logrus"
)

// App serves api
type App struct {
	logrus.FieldLogger
	data.Repo
}

// NewApp return app instance
func NewApp(logger logrus.FieldLogger, db *sql.DB, redis *redis.Client) *App {
	response := shared.NewResponse(logger)
	repo := data.NewRepo(logger, db, response, redis)
	return &App{
		logger,
		repo,
	}
}
