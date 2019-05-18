package cmd

import (
	"database/sql"
	"mess-app/internal/api"
	"mess-app/shared"
	"net/http"
	"os"
)

func serveApp(db *sql.DB, c *configs) {
	done := make(chan int)
	logger := shared.NewLogger(c.logger.level, os.Stdout)
	app := api.NewApp(logger, db)
	go func() {
		err := http.ListenAndServe(c.appConfig.host+":"+c.appConfig.port, app.Router())
		if err != nil {
			logger.Error(err)
		}
		done <- 1
	}()
	<-done
}
