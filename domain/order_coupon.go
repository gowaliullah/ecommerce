package domain

type OrderCoupon struct {
	ID              int     `json:"id" db:"id"`
	OrderID         int     `json:"order_id" db:"order_id"`
	CouponID        int     `json:"coupon_id" db:"coupon_id"`
	DiscountApplied float64 `json:"discount_applied" db:"discount_applied"`
}
