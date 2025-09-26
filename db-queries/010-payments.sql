

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    amount NUMERIC(10,2) NOT NULL,
    provider VARCHAR(50), -- stripe, paypal etc.
    status VARCHAR(20) DEFAULT 'pending', -- pending, success, failed
    transaction_id VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



DROP TABLE payments;