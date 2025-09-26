package domain

import "time"

type Address struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Line1     string    `json:"line1" db:"line1"`
	Line2     *string   `json:"line2,omitempty" db:"line2"`
	City      string    `json:"city" db:"city"`
	State     *string   `json:"state,omitempty" db:"state"`
	Zip       string    `json:"zip" db:"zip"`
	Country   string    `json:"country" db:"country"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
