package model

import "time"

type ItemDetail struct {
	ID       int      `json:"id"`
	ItemCode string   `json:"item_code"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Location Location `json:"location"`
	Price    int      `json:"price"` // Consider using int64 if needed
	Stock    int      `json:"stock"` // Consider using int64 if needed
}

// Transaction represents a record of a transaction
type Transaction struct {
	ID           int        `json:"id"`
	ProductID    int        `json:"product_id"`
	Product      ItemDetail `json:"product"` // Tambahkan field ini
	Quantity     int        `json:"quantity"`
	MovementType string     `json:"movement_type"`
	Timestamp    time.Time  `json:"timestamp"`
}
