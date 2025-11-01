package domain

import "time"

type Cart struct {
	ID        int       `json:"id" db:"id"`
	Uuid      string    `json:"unique_id" db:"unique_id" validate:"required,unique"`
	UserID    *string   `json:"user_id" db:"user_id"`
	ProductID string    `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
