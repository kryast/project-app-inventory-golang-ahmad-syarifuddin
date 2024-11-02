package model

import "time"

type InOutProduct struct {
	ID           int       `json:"id"`
	ProductID    int       `json:"product_id"`
	Quantity     int       `json:"quantity"`
	MovementType string    `json:"movement_type"`
	Timestamp    time.Time `json:"timestamp"` // Add this field
}
