package main

import (
	"log"
	"net/http"

	"github.com/CBYeuler/go-rest-crud-api/db"
	"github.com/CBYeuler/go-rest-crud-api/routes"
)

func main() {
	db.Init()
	// Initialize the database connection

	log.Println("Database initialized successfully")

	// Set up the router and start the server

	router := routes.SetupRouters()

	log.Println("Starting server on :8080")

	// Start the HTTP server on port 8080
	// The server will listen for incoming requests and route them accordingly

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server stopped")
}

// The main function initializes the database and starts the HTTP server
// It sets up the router to handle incoming requests and routes them to the appropriate handlers
