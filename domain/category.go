package domain

import "time"

type Category struct {
	ID        int       `json:"id" db:"id"`
	Uuid      string    `json:"unique_id" db:"unique_id" validate:"required,unique"`
	Name      string    `json:"name" db:"name"`
	Slug      string    `json:"slug" db:"slug"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
