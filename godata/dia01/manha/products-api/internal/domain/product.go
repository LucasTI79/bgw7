package domain

// Product is a struct that contains the information of a product
type Product struct {
	Id       int
	Name     string
	Type     string
	Quantity int
	Price    float64
}

type ProductAttributes struct {
	Name     string
	Type     string
	Quantity int
	Price    float64
}

func (pa ProductAttributes) ToDomain() *Product {
	return &Product{
		Name:     pa.Name,
		Type:     pa.Type,
		Quantity: pa.Quantity,
		Price:    pa.Price,
	}
}

type ProductRepository interface {
	// Get returns all the products
	Get() ([]Product, error)
	// GetByID returns a product by id
	GetByID(productId int) (*Product, error)
	// Save saves a product
	Save(product *Product) error
	// Update updates a product
	Update(product *Product) error
	// Delete a product
	Delete(productId int) error
}
