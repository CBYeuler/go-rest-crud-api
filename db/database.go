package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// Package db provides database initialization and connection management
// It uses SQLite as the database engine

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "./items.db") // Open a connection to the SQLite database
	// The database file is named items.db and is located in the current directory
	if err != nil {
		log.Fatal("Failed to connect to database:", err) // Initialize the database connection
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
