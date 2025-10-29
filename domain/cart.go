package domain

import "time"

type Cart struct {
	ID        int       `json:"id" db:"id"`
	Uuid      string    `json:"unique_id" db:"unique_id" validate:"required,unique"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
