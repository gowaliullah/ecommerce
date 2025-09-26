package domain

import "time"

type Payment struct {
	ID            int       `json:"id" db:"id"`
	OrderID       int       `json:"order_id" db:"order_id"`
	Amount        float64   `json:"amount" db:"amount"`
	Provider      string    `json:"provider" db:"provider"`
	Status        string    `json:"status" db:"status"`
	TransactionID string    `json:"transaction_id" db:"transaction_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
