package models

import "time"

// Product represents a product in the system.

type Product struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Version   string    `db:"version" json:"version"`
	EOLDate   time.Time `db:"eol_date" json:"eol_date"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
