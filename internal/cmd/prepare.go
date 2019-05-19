package cmd

import (
	"database/sql"
	"fmt"
	"mess-app/internal/core"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"

	// postgres initialization
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func mustPrepareDB() (*sql.DB, error) {
	pqstring := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("app.db.host"),
		viper.GetInt("app.db.port"),
		viper.GetString("app.db.user"),
		viper.GetString("app.db.password"),
		viper.GetString("app.db.dbname"),
	)
	db, err := sql.Open("postgres", pqstring)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db connection")
	}
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "Cannot ping to db")
	}
	return db, nil
}

func mustPrepareRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("app.redis.host"), viper.GetString("app.redis.port")),
		Password: viper.GetString("app.redis.password"),
		DB:       0,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.Wrap(err, "Failed to prepare redis")
	}
	return client, nil
}

func mustPrepareLogger() (logrus.FieldLogger, error) {
	var level logrus.Level
	var out *os.File
	switch viper.GetString("app.logger.level") {
	case core.LOG_DEBUG:
		level = logrus.DebugLevel
	case core.LOG_INFO:
		level = logrus.InfoLevel
	case core.LOG_WARNING:
		level = logrus.WarnLevel
	case core.LOG_ERROR:
		level = logrus.ErrorLevel
	default:
		return nil, errors.New("Invalid log level")
	}
	switch viper.GetString("app.logger.output") {
	case core.OUT_STDOUT:
		out = os.Stdout
	default:
		return nil, errors.New("Invalid log output")
	}

	l := &logrus.Logger{
		Out:       out,
		Level:     level,
		Formatter: &logrus.JSONFormatter{},
	}
	return l, nil
}
