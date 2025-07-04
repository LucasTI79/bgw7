package product

const (
	GetQuery         = "SELECT product_id, name, type, quantity, price FROM products"
	GetOneQuery      = "SELECT product_id, name, type, quantity, price FROM products WHERE product_id = ?"
	StoreQuery       = "INSERT INTO products (name, type, quantity, price) VALUES (?, ?, ?, ?)"
	UpdateQuery      = "UPDATE products SET name=?, type=?, quantity=?, price=? WHERE product_id=?"
	DeleteQuery      = "DELETE FROM products WHERE product_id=?"
	GetWithInfoQuery = `
		SELECT 
			p.product_id, 
			p.name, 
			p.type, 
			p.quantity, 
			p.price,
			w.name AS warehouseName,
			w.address
		FROM products AS p
		INNER JOIN warehouses AS w 
			ON w.warehouse_id = p.warehouse_id;
	`
)
