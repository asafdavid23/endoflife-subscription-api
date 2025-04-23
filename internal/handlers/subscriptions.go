package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/asafdavid23/endoflife-subscription-api/internal/db"
)

type CreateSubscriptionRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Product          string `json:"product"`
	Version          string `json:"version"`
	EOLDate          string `json:"eol_date"`
	NotifyDaysBefore int    `json:"notify_days_before"`
}

func CreateSubscriptionHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateSubscriptionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		eolDate, err := time.Parse("2006-01-02", req.EOLDate)
		if err != nil {
			http.Error(w, "Invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}

		sub := &db.Subscription{
			Name:             req.Name,
			Email:            req.Email,
			Product:          req.Product,
			Version:          req.Version,
			EOLDate:          eolDate,
			NotifyDaysBefore: req.NotifyDaysBefore,
			Notified:         false,
		}

		if err := db.CreateSubscription(dbConn, sub); err != nil {
			http.Error(w, "Failed to create subscription", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Subscription created successfully"})
	}
}

func ListSubscriptionsHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subs, err := db.GetAllSubscriptions(dbConn)
		if err != nil {
			http.Error(w, "Failed to fetch subscriptions", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subs)
	}
}
