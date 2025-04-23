package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"

	"endoflife-subscription-api/internal/db"

	"github.com/asafdavid23/endoflife-subscription-api/internal/handler"
)

func main() {
	dbConn, err := sql.Open("sqlite3", "file:subscriptions.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer dbConn.Close()

	if err := db.InitializeSchema(dbConn); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/subscriptions", handler.CreateSubscriptionHandler(dbConn))
		r.Get("/subscriptions", handler.ListSubscriptionsHandler(dbConn))
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
