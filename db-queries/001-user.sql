-- Active: 1755325089603@@127.0.0.1@5432@ecommerce


CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    -- role user_role DEFAULT 'user',
    role VARCHAR(25) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


DROP ENUM user_role;

DROP TABLE users;


INSERT INTO users (first_name, last_name, email, password, role, created_at, updated_at) VALUES
('Alice', 'Williams', 'alice.williams@example.com', 'hashed_password_001', 'admin', '2025-10-07 10:00:00', '2025-10-07 10:00:00'),
('Bob', 'Jones', 'bob.jones@example.com', 'hashed_password_002', 'user', '2025-10-07 11:15:00', '2025-10-07 11:15:00'),
('Carol', 'Davis', 'carol.davis@example.com', 'hashed_password_003', 'user', '2025-10-07 12:00:00', '2025-10-07 12:22:00'),
('David', 'Clark', 'david.clark@example.com', 'hashed_password_004', 'user', '2025-10-07 09:30:00', '2025-10-07 10:45:00'),
('Eve', 'Martinez', 'eve.martinez@example.com', 'hashed_password_005', 'user', '2025-10-07 08:20:00', '2025-10-07 08:20:00');

