package db

import "database/sql"

func InitializeSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		product TEXT NOT NULL,
		version TEXT NOT NULL,
		eol_date DATE NOT NULL,
		notify_days_before INTEGER NOT NULL,
		notified BOOLEAN NOT NULL DEFAULT 0
	);`
	_, err := db.Exec(schema)
	return err
}
