package domain

import "time"

type CartItem struct {
	ID        int       `json:"id" db:"id"`
	Uuid      string    `json:"unique_id" db:"unique_id" validate:"required,unique"`
	CartID    string    `json:"cart_id" db:"cart_id"`
	ProductID string    `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
