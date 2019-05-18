package cmd

import (
	"database/sql"
	"fmt"

	// postgres initialization
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func mustPrepareDB(c *configs) (*sql.DB, error) {
	pqstring := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.dbConfig.host,
		c.dbConfig.port,
		c.dbConfig.user,
		c.dbConfig.password,
		c.dbConfig.dbName,
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
