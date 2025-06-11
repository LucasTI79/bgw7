package output

import "github.com/bgw7/products-api/internal/domain"

type CreateBodyProductOutput struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}

type UpdateBodyProductOutput struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}

type GetBodyProductOutput struct {
	Message string           `json:"message"`
	Data    []domain.Product `json:"data"`
	Error   bool             `json:"error"`
}

type GetOneBodyProductOutput struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}
