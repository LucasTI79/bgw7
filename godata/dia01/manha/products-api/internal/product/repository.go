package product

import (
	"database/sql"

	"github.com/bgw7/products-api/internal/domain"
)

func NewRepository(db *sql.DB) domain.ProductRepository {
	return &mysqlRepository{
		db: db,
	}
}
