package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func PostgresDBConnection() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("Error: DATABASE_URL is not set in environment variables")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error: Not connected to the database. %w", err)
	}
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("Error: Unable to ping the database. %w", err)
	}
	return db, nil
}