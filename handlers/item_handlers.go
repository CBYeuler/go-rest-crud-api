package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CBYeuler/go-rest-crud-api/db"
	"github.com/CBYeuler/go-rest-crud-api/models"
	"github.com/gorilla/mux"
)

// CreateItem handles POST /items
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	// Decode request body into item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Close the request body after reading
	defer r.Body.Close()
	// Ensure the database is initialized
	if db.DB == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}

	// Validate input
	if item.Name == "" || item.Price <= 0 {
		http.Error(w, "Invalid item data", http.StatusBadRequest)
		return
	}

	// Prepare SQL statement
	stmt, err := db.DB.Prepare("INSERT INTO items(name, price) VALUES(?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the statement
	result, err := stmt.Exec(item.Name, item.Price)
	if err != nil {
		http.Error(w, "Failed to execute insert", http.StatusInternalServerError)
		return
	}

	// Get the new item ID
	id, _ := result.LastInsertId()
	item.ID = int(id)

	// Respond with the created item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, price FROM items")
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			http.Error(w, "Failed to scan item", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := db.DB.QueryRow("SELECT id, name, price FROM items WHERE id = ?", id)

	var item models.Item
	if err := row.Scan(&item.ID, &item.Name, &item.Price); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch item", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if item.Name == "" || item.Price <= 0 {
		http.Error(w, "Invalid item data", http.StatusBadRequest)
		return
	}

	stmt, err := db.DB.Prepare("UPDATE items SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.Name, item.Price, id)
	if err != nil {
		http.Error(w, "Failed to execute update", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	// Convert id from string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	item.ID = intID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	stmt, err := db.DB.Prepare("DELETE FROM items WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		http.Error(w, "Failed to execute delete", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // No content to return
}

// The above code defines the handlers for CRUD operations on items
// It includes functions to create, read, update, and delete items in the database
// Each function handles the HTTP request, interacts with the database, and returns the appropriate response
// The CreateItem function handles POST requests to create a new item
// The GetItems function handles GET requests to retrieve all items
// The GetItem function handles GET requests to retrieve a specific item by ID
// The UpdateItem function handles PUT requests to update an existing item by ID
// The DeleteItem function handles DELETE requests to remove an item by ID
// The handlers use the database connection from the db package to interact with the SQLite database
// The functions handle errors appropriately and return JSON responses
