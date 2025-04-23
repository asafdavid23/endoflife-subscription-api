package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "file:subscriptions.db?cache=shared&mode=rwc")

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	defer db.Close()

	if err := InitializeSchema(db); err != nil {
		log.Fatalf("failed to initialize schema: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("	/api/v1/subscriptions", func(r chi.Router) {
		r.Post("/subscription", CraeteSubscriptionHandler(db))
		r.Get("/subscriptions", GetAllSubscriptionsHandler(db))
	})

	log.Println("Starting server on :8080")

	http.ListenAndServe(":8080", r)
}
