package db

import (
	"database/sql"
	"log"

	//"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./items.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS items (
	id INTEGER PRIMARY KEY AUTOINCREMENT,	
	name TEXT NOT NULL,
	price REAL NOT NULL,
	
);`
	_, err = DB.Exec(createTable)

	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

}
