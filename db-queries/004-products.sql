

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    stock INT DEFAULT 0,
    img_url VARCHAR(255) NOT NULL,
    -- category_id INT REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

 
DROP TABLE products;




INSERT INTO products (title, description, price, stock, img_url, created_at, updated_at) VALUES
('Wireless Mouse', 'Ergonomic wireless mouse with adjustable DPI settings.', 29.99, 150, 'https://example.com/images/mouse.jpg', '2025-10-01 10:00:00', '2025-10-01 10:00:00'),
('Bluetooth Headphones', 'Noise-cancelling over-ear headphones with 20-hour battery life.', 89.99, 75, 'https://example.com/images/headphones.jpg', '2025-10-02 12:30:00', '2025-10-03 09:15:00'),
('USB-C Hub', 'Multi-port USB-C hub with HDMI, USB 3.0, and SD card reader.', 45.50, 200, 'https://example.com/images/usb_hub.jpg', '2025-10-03 14:20:00', '2025-10-04 11:10:00'),
('Gaming Keyboard', 'Mechanical keyboard with RGB lighting and programmable keys.', 129.99, 50, 'https://example.com/images/keyboard.jpg', '2025-10-04 08:45:00', '2025-10-05 16:25:00'),
('Smartwatch', 'Fitness tracker with heart rate monitor and GPS.', 199.99, 30, 'https://example.com/images/smartwatch.jpg', '2025-10-05 09:00:00', '2025-10-06 13:40:00');



SELECT * FROM products;