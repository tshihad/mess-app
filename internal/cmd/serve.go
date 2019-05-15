package cmd

import (
	"database/sql"
	"fmt"
	"mess-app/internal/api"
	"mess-app/shared"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func serveApp(db *sql.DB) {
	done := make(chan int)
	logger := shared.NewLogger(logrus.DebugLevel, os.Stdout)
	app := api.NewApp(logger, db)
	go func() {
		err := http.ListenAndServe(":8080", app.Router())
		if err != nil {
			fmt.Println(err)
		}
		done <- 1
	}()
	<-done
}
