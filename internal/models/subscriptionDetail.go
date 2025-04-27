package models

import "time"

type SubscriptionDetail struct {
	SubscriptionID   int       `db:"subscription_id" json:"subscription_id"`
	UserEmail        string    `db:"user_email" json:"user_email"`
	ProductName      string    `db:"product_name" json:"product_name"`
	ProductVersion   string    `db:"product_version" json:"product_version"`
	ProductEOLDate   time.Time `db:"product_eol_date" json:"product_eol_date"`
	NotifyBeforeDays int       `db:"notify_before_days" json:"notify_before_days"`
	SubscribedAt     time.Time `db:"created_at" json:"subscribed_at"`
}
