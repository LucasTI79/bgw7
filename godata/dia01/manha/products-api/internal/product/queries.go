package product

const (
	GetQuery    = "SELECT product_id, name, type, quantity, price FROM products"
	GetOneQuery = "SELECT product_id, name, type, quantity, price FROM products WHERE product_id = ?"
	StoreQuery  = "INSERT INTO products (name, type, quantity, price) VALUES (?, ?, ?, ?)"
	UpdateQuery = "UPDATE products SET name=?, type=?, quantity=?, price=? WHERE product_id=?"
	DeleteQuery = "DELETE FROM products WHERE product_id=?"
)
