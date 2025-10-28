-- Active: 1755325089603@@127.0.0.1@5432@ecommerce


CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE users (
    id SERIAL NOT NULL,
    unique_id UUID UNIQUE PRIMARY KEY NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(25) DEFAULT 'user', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- role user_role DEFAULT 'user',
);


DROP ENUM user_role;

DROP TABLE users;


INSERT INTO users (id, unique_id, first_name, last_name, email, password, role, created_at, updated_at) VALUES
(1, '550e8400-e29b-41d4-a716-446655440000', 'Alice', 'Smith', 'smith@example.com', 'hashed_password_123', 'admin', '2025-10-01 10:00:00', '2025-10-01 10:00:00'),
(2, '6ba7b810-9dad-11d1-80b4-00c04fd430c8', 'Bob', 'Johnson', 'bob.johnson@example.com', 'hashed_password_456', 'user', '2025-10-02 12:30:00', '2025-10-02 12:30:00'),
(3, '7d793037-a076-4b61-8e4d-7e4b7e4b7e4b', 'Clara', 'Davis', 'clara.davis@example.com', 'hashed_password_789', 'editor', '2025-10-03 15:45:00', '2025-10-04 09:15:00'),
(4, '8f8b8f8b-8f8b-4b61-8e4d-8f8b8f8b8f8b', 'David', 'Brown', 'david.brown@example.com', 'hashed_password_012', 'user', '2025-10-04 08:20:00', '2025-10-04 08:20:00'),
(5, '9c9c9c9c-9c9c-4b61-8e4d-9c9c9c9c9c9c', 'Emma', 'Wilson', 'emma.wilson@example.com', 'hashed_password_345', 'moderator', '2025-10-05 14:10:00', '2025-10-05 14:10:00');