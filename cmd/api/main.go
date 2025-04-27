package main

import (
	"log"
	"net/http"

	"github.com/asafdavid23/eol-api/internal/handlers"
	"github.com/asafdavid23/eol-api/pkg/db"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Simple health check
	r.Get("/healthz", handlers.HealthHandler)

	// Future: you can add routes like:
	// r.Post("/subscribe", handlers.SubscribeHandler(db))

	return r
}

func main() {
	// Connect to the database
	db, err := db.DBConnect()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Close the database connection when the main function exits
	defer db.Close()

	// Setup router
	r := setupRouter()

	port := "8080"

	log.Printf("Starting server on port %s", port)

	// Start the server
	err = http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
