CREATE DATABASE IF NOT EXISTS orders;

USE orders;

CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY,
    price DECIMAL(10,2),
    tax DECIMAL(10,2),
    final_price DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO orders (id, price, tax, final_price, created_at, updated_at) 
VALUES 
(1, 100.50, 10.05, 110.55, NOW(), NOW()),
(2, 250.75, 25.08, 275.83, NOW(), NOW()),
(3, 89.99, 8.99, 98.98, NOW(), NOW()),
(4, 150.00, 15.00, 165.00, NOW(), NOW()),
(5, 500.25, 50.02, 550.27, NOW(), NOW());
