package domain

import "time"

type Coupon struct {
	ID            int       `json:"id" db:"id"`
	Code          string    `json:"code" db:"code"`
	DiscountType  string    `json:"discount_type" db:"discount_type"`
	DiscountValue float64   `json:"discount_value" db:"discount_value"`
	ExpiryDate    time.Time `json:"expiry_date" db:"expiry_date"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
