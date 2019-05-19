package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"mess-app/internal/api"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/go-redis/redis"
)

func serveApp(db *sql.DB, redis *redis.Client, logger logrus.FieldLogger) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
		redis.Close()
		db.Close()
	}()
	done := make(chan int)
	app := api.NewApp(logger, db, redis)
	go func() {
		err := http.ListenAndServe(
			fmt.Sprintf("%s:%d", viper.GetString("app.appconfig.host"), viper.GetInt("app.appconfig.port")),
			app.Router(),
		)
		if err != nil {
			logger.Error(err)
		}
		done <- 1
	}()
	<-done
}
