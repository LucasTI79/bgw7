package product

import (
	"fmt"

	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/pkg/apperrors"
)

type mapRepository struct {
	storage map[int]domain.ProductAttributes
	lastId  int
}

func (r mapRepository) Get() ([]domain.Product, error) {
	products := make([]domain.Product, 0)

	for productId, productAttr := range r.storage {
		product := productAttr.ToDomain()
		product.Id = productId
		products = append(products, *product)
	}

	return products, nil
}

func (r mapRepository) GetByID(productId int) (*domain.Product, error) {
	p, ok := r.storage[productId]

	if !ok {
		return nil, apperrors.ErrResourceNotFound
	}

	product := p.ToDomain()
	product.Id = productId

	return product, nil
}

func (r *mapRepository) Save(product *domain.Product) error {

	attr := domain.ProductAttributes{
		Name:     product.Name,
		Type:     product.Type,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	r.lastId++
	r.storage[r.lastId] = attr
	product.Id = r.lastId
	return nil
}

func (r *mapRepository) Update(product *domain.Product) error {
	attr := domain.ProductAttributes{
		Name:     product.Name,
		Type:     product.Type,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
	// update
	_, ok := r.storage[product.Id]

	if !ok {
		return apperrors.ErrResourceNotFound
	}

	r.storage[product.Id] = attr
	return nil
}

func (r *mapRepository) Delete(productId int) error {
	if _, ok := r.storage[productId]; !ok {
		return fmt.Errorf("%w: product with %d not found", apperrors.ErrResourceNotFound, productId)
	}
	// delete
	delete(r.storage, productId)
	return nil

}

func (r mapRepository) GetWithInfo() ([]domain.ProductInfo, error) {
	products := make([]domain.ProductInfo, 0)

	for productId, productAttr := range r.storage {
		product := productAttr.ToDomain()
		product.Id = productId

		productInfo := domain.ProductInfo{
			Product: *product,
		}

		products = append(products, productInfo)
	}
	return products, nil
}
