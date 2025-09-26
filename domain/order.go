package domain

import "time"

// OrderStatus type for ENUM mapping
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusDelivered OrderStatus = "delivered"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID        int         `json:"id" db:"id"`
	UserID    int         `json:"user_id" db:"user_id"`
	Total     float64     `json:"total_price" db:"total_price"`
	Status    OrderStatus `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}
