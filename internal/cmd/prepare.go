package cmd

import (
	"database/sql"
	"fmt"

	// postgres initialization
	_ "github.com/lib/pq"
)

func mustPrepareDB() (*sql.DB, error) {
	conf := make(map[string]string)
	conf["host"] = "localhost"
	conf["dbname"] = "mess"
	conf["password"] = "pswd123"
	conf["port"] = "5432"
	conf["user"] = "postgres"

	pqstring := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf["host"], conf["port"], conf["user"], conf["password"], conf["dbname"])
	return sql.Open("postgres", pqstring)
}
