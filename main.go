package main

import (
	"log"
	"net/http"

	"github.com/CBYeuler/go-rest-crud-api/db"
	"github.com/CBYeuler/go-rest-crud-api/routes"
)

func main() {
	db.Init()

	router := routes.SetupRouter()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server stopped")
}
