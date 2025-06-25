package routes

import (

	// Import the handlers package for route handling
	"github.com/CBYeuler/go-rest-crud-api/handlers"
	"github.com/gorilla/mux"
)

// Package routes defines the routes for the application
// It uses the Gorilla Mux router to handle HTTP requests
// It imports the handlers package to route requests to the appropriate handler functions
// It provides a function to set up the router and define the routes
// SetupRouter initializes the router and defines the routes for the application

func SetupRouters() *mux.Router {
	router := mux.NewRouter()

	// Define the routes for the application
	router.HandleFunc("/items", handlers.GetItem).Methods("GET")                   // Get all items
	router.HandleFunc("/items/{id:[0-9]+}", handlers.GetItem).Methods("GET")       // Get a specific item by ID
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")               // Create a new item
	router.HandleFunc("/items/{id:[0-9]+}", handlers.UpdateItem).Methods("PUT")    // Update an existing item by ID
	router.HandleFunc("/items/{id:[0-9]+}", handlers.DeleteItem).Methods("DELETE") // Delete an item by ID

	return router
}
