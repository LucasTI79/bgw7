package main

import (
	"errors"
	"fmt"
	"net/http"
)

type Product struct {
	Name        string
	ProductCode string
}

// DI -> inversão de dependência

type Repository struct{}

func (r Repository) Save(product Product) error {
	return ErrProductCodeAlreadyExists
}

type Service struct {
	Repository
}

func (s Service) Save(product Product) error {
	err := s.Repository.Save(product)

	return fmt.Errorf("%w: %s", err, product.ProductCode)
}

type Handler struct {
	Service
}

func (h Handler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// 400 - Bad request
// 401 - Unauthorized
// 403 - Forbidden
// 404 - Not found
// 409 - Conflict
// 422 - Unprocessable Entity
// 429 - Too many requests

var ErrProductCodeAlreadyExists = errors.New("product code already exists")
var ErrProductNameIsRequired = errors.New("product name is required")

func main() {
	product := Product{}

	repo := Repository{}
	service := Service{repo}

	err := service.Save(product)

	switch {
	case errors.Is(err, ErrProductNameIsRequired):
		fmt.Println(http.StatusBadRequest)
	case errors.Is(err, ErrProductCodeAlreadyExists):
		fmt.Println(http.StatusConflict)
	default:
		fmt.Println(http.StatusInternalServerError)
	}
}
