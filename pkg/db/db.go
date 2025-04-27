package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DBConnect() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=yourpassword dbname=yourdb sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
