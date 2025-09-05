CREATE TABLE IF NOT EXISTS products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL, 
    type VARCHAR(50) NOT NULL,
    quantity INT DEFAULT 0,
    price DECIMAL(12,2) DEFAULT 0
);

-- INSERT INTO products(name, type, quantity, price) 
-- VALUES 
-- ('Batata', 'vegetal', 1, 6.00),
-- ('Cenoura', 'vegetal', 1, 4.00);

-- SELECT * FROM products;

ALTER TABLE products ADD COLUMN warehouse_id INT;

CREATE TABLE IF NOT EXISTS warehouses(
    warehouse_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    address VARCHAR(100) NOT NULL
);

ALTER TABLE products ADD CONSTRAINT FOREIGN KEY(warehouse_id) REFERENCES warehouses(warehouse_id);

-- DESC products;
-- GMUD

-- INSERT INTO warehouses(name, address) 
-- VALUES 
-- ('CD pataxos', 'pataxos'),
-- ('CD cajamar', 'cajamar');

-- UPDATE products SET warehouse_id = 1 WHERE product_id = 1;
-- UPDATE products SET warehouse_id = 2 WHERE product_id = 2;


-- SELECT 
-- 	p.product_id, 
--   p.name, 
--   p.type, 
--   p.quantity, 
--   p.price,
--   w.name AS warehouseName,
--   w.address
-- FROM products AS p
-- INNER JOIN warehouses AS w 
-- 	ON w.warehouse_id = p.warehouse_id;
  