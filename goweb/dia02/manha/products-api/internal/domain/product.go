package domain

// Product is a struct that contains the information of a product
type Product struct {
	Id       int
	Name     string
	Type     string
	Quantity int
	Price    float64
}

type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Mapper
func (r RequestBodyProduct) ToDomain() Product {
	return Product{
		Name:     r.Name,
		Type:     r.Type,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}

type ResponseBodyProduct struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
	Error   bool     `json:"error"`
}
