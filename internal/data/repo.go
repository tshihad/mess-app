package data

import (
	"database/sql"
	"mess-app/shared"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// TODO add cache
type repo struct {
	logrus.FieldLogger
	*sql.DB
	shared.Response
	Cache
}

// NewRepo returns repo instance
func NewRepo(logger logrus.FieldLogger, db *sql.DB, response shared.Response, redis *redis.Client) *repo {
	return &repo{
		FieldLogger: logger,
		DB:          db,
		Response:    response,
		Cache: &redisCache{
			c: redis,
		},
	}
}
