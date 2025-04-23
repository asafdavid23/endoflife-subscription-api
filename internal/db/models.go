// Models
package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Subscription struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Product          string    `json:"product"`
	Version          string    `json:"version"`
	EOLDate          time.Time `json:"eol_date"`
	NotifyDaysBefore int       `json:"notify_days_before"`
	Notified         bool      `json:"notified"`
}

func CraeteSubscription(db *sql.DB, sub *Subscription) error {
	query := `INSERT INTO subscriptions (name, email, product, version, eol_date, notify_days_before, notified) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query, sub.Name, sub.Email, sub.Product, sub.Version, sub.EOLDate, sub.NotifyDaysBefore, sub.Notified)

	if err != nil {
		return fmt.Errorf("failed to create subscription: %v", err)
	}

	return nil
}

func GetAllSubscriptions(db *sql.DB) ([]Subscription, error) {
	rows, err := db.Query("SELECT id, name, email, product, version, eol_date, notify_days_before, notified FROM subscriptions")

	if err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %v", err)
	}

	defer rows.Close()

	subs := []Subscription{}

	for rows.Next() {
		s := Subscription{}

		err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Product, &s.Version, &s.EOLDate, &s.NotifyDaysBefore, &s.Notified)

		if err != nil {
			return nil, fmt.Errorf("failed to scan subscription: %v", err)
		}

		subs = append(subs, s)
	}

	return subs, nil
}
