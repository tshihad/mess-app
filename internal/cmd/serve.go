package cmd

import (
	"database/sql"
	"fmt"
	"mess-app/internal/api"
	"net/http"
)

func serveApp() {
	done := make(chan int)
	app := api.NewApp(nil, &sql.DB{})
	go func() {
		err := http.ListenAndServe(":8080", app.Router())
		if err != nil {
			fmt.Println(err)
		}
		done <- 1
	}()
	<-done
}
