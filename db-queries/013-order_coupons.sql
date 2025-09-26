

CREATE TABLE order_coupons (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    coupon_id INT REFERENCES coupons(id) ON DELETE SET NULL,
    discount_applied NUMERIC(10,2) NOT NULL
);



DROP TABLE order_coupons;