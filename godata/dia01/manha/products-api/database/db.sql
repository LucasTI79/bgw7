-- Id       int
-- Name     string
-- Type     string
-- Quantity int
-- Price    float64


CREATE TABLE IF NOT EXISTS products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL, 
    type VARCHAR(50) NOT NULL,
    quantity INT DEFAULT 0,
    price DECIMAL(12,2) DEFAULT 0
);