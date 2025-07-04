package mocks

import (
	"github.com/bgw7/products-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r *ProductRepositoryMock) Get() ([]domain.Product, error) {
	args := r.Mock.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (r *ProductRepositoryMock) GetByID(productId int) (*domain.Product, error) {
	args := r.Mock.Called(productId)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (r *ProductRepositoryMock) Save(product *domain.Product) error {
	args := r.Mock.Called(product)
	return args.Error(0)
}

func (r *ProductRepositoryMock) Update(product *domain.Product) error {
	args := r.Mock.Called(product)
	return args.Error(0)
}

func (r *ProductRepositoryMock) Delete(productId int) error {
	args := r.Mock.Called(productId)
	return args.Error(0)
}
