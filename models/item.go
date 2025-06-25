package models

type Item struct {
	ID    int     `json:"id"`    // Unique identifier for the item
	Name  string  `json:"name"`  // Name of the item
	Price float64 `json:"price"` // Price of the item
}

// TableName returns the name of the table for the Item model
