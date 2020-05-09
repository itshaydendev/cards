package internal

import (
	"database/sql"

	"github.com/itshaydendev/cards/pkg/logger"

	// The PostgreSQL database driver
	_ "github.com/lib/pq"
)

// Database creates the connection to Postgres and gives access to it
func Database() (*sql.DB, error) {
	connStr := "postgres://cards:cards@localhost:5432/cards?sslmode=disable"
	var err error

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		logger.Error("Failed to connect to database.")
		logger.Error(err.Error())
		return nil, err
	}

	return db, nil
}
