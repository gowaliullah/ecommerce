-- Active: 1755325089603@@127.0.0.1@5432@laundry

CREATE TABLE shops (
    id SERIAL PRIMARY KEY,
    shop_name VARCHAR(150),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   
);


INSERT INTO shops (shop_name)
VALUES
('Tech Haven'),
('Green Valley Organics'),
('FitLife Sports'),
('Modern Home Essentials'),
('Pixel Planet Gaming');


DROP TABLE shops;