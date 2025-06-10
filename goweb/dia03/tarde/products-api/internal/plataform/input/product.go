package input

import "github.com/bgw7/products-api/internal/domain"

type CreateBodyProductInput struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Mapper
func (r CreateBodyProductInput) ToDomain() domain.Product {
	return domain.Product{
		Name:     r.Name,
		Type:     r.Type,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}

type UpdateBodyProductInput struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Mapper
func (r UpdateBodyProductInput) ToDomain() domain.Product {
	return domain.Product{
		Name:     r.Name,
		Type:     r.Type,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}
