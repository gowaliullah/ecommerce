-- +migrate Up


CREATE TABLE IF NOT EXISTS cartitems (
    id SERIAL PRIMARY KEY,
    unique_id UUID DEFAULT gen_random_uuid() UNIQUE,
    user_id UUID REFERENCES users(unique_id) ON DELETE SET NULL, 
    product_id UUID REFERENCES products(unique_id) ON DELETE SET NULL, 
    quantity INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);  