package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Sahil2004/gmarket/server/config"
	_ "github.com/lib/pq"
)

func PostgresDBConnection() (*sql.DB, error) {
	connStr := config.AppConfig().DatabaseURL
	if connStr == "" {
		return nil, fmt.Errorf("Error: DATABASE_URL is not set in environment variables")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error: Not connected to the database. %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("Error: Unable to ping the database. %w", err)
	}
	return db, nil
}
