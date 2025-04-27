package models

import "time"

type Subscription struct {
	ID               int       `db:"id" json:"id"`
	UserID           int       `db:"user_id" json:"user_id"`
	ProductID        int       `db:"product_id" json:"product_id"`
	NotifyBeforeDays int       `db:"notify_before_days" json:"notify_before_days"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
