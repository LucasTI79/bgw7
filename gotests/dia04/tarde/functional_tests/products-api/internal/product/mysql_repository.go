package product

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/pkg/apperrors"
)

type mysqlRepository struct {
	db *sql.DB
}

// Get implements domain.ProductRepository.
func (m *mysqlRepository) Get() ([]domain.Product, error) {
	rows, err := m.db.Query(GetQuery)
	if err != nil {
		return nil, err
	}

	var result []domain.Product
	defer rows.Close()
	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Quantity, &product.Price); err != nil {
			return nil, err
		}
		result = append(result, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// GetByID implements domain.ProductRepository.
func (m *mysqlRepository) GetByID(productId int) (*domain.Product, error) {
	row := m.db.QueryRow(GetOneQuery, productId)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Quantity, &product.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[%w] product with id: %d don't exists", apperrors.ErrResourceNotFound, productId)
		}
		return nil, err
	}

	return &product, nil
}

// Save implements domain.ProductRepository.
func (m *mysqlRepository) Save(product *domain.Product) error {
	// query execution
	result, err := m.db.Exec(
		StoreQuery,
		(*product).Name, (*product).Type, (*product).Quantity, (*product).Price,
	)

	if err != nil {
		return err
	}

	// get last inserted id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	// set user id
	(*product).Id = int(lastInsertId)
	return nil
}

// Update implements domain.ProductRepository.
func (m *mysqlRepository) Update(product *domain.Product) error {
	// query execution
	_, err := m.db.Exec(
		UpdateQuery,
		(*product).Name, (*product).Type, (*product).Quantity, (*product).Price, (*product).Id,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.ProductRepository.
func (m *mysqlRepository) Delete(productId int) error {
	_, err := m.db.Exec(DeleteQuery, productId)
	if err != nil {
		return err
	}
	return nil
}

// GetWithInfo implements domain.ProductRepository.
func (m *mysqlRepository) GetWithInfo() ([]domain.ProductInfo, error) {
	rows, err := m.db.Query(GetWithInfoQuery)
	if err != nil {
		return nil, err
	}

	var result []domain.ProductInfo
	defer rows.Close()
	for rows.Next() {
		var productInfo domain.ProductInfo

		if err := rows.Scan(
			&productInfo.Id,
			&productInfo.Name,
			&productInfo.Type,
			&productInfo.Quantity,
			&productInfo.Price,
			&productInfo.WarehouseName,
			&productInfo.WarehouseAddress,
		); err != nil {
			return nil, err
		}
		result = append(result, productInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
