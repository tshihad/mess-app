package cmd

import (
	"database/sql"
	"fmt"

	// postgres initialization
	_ "github.com/lib/pq"
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
	return sql.Open("postgres", pqstring)
}
