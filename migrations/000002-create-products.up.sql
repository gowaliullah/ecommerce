-- +migrate Up

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    unique_id UUID DEFAULT gen_random_uuid() UNIQUE,      
    title VARCHAR(255) NOT NULL, 
    description TEXT,  
    price NUMERIC(10,2) NOT NULL,  
    stock INT DEFAULT 0,
    img_url VARCHAR(255) NOT NULL,
    created_by UUID REFERENCES users(unique_id) ON DELETE SET NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    -- category_id INT REFERENCES categories(id) ON DELETE SET NULL
);
