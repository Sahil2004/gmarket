package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PostgresDBConnection() (*sql.DB, error) {
	connStr := "user=sahil password=postgres dbname=gmarket sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error: Not connected to the database. %w", err)
	}
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("Error: Unable to ping the database. %w", err)
	}
	defer db.Close()
	return db, nil
}