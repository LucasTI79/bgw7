package product_test

import (
	"testing"

	"github.com/bgw7/products-api/config"
	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/internal/product"
	"github.com/bgw7/products-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestProductStore(t *testing.T) {
	config.Init()
	utils.RegisterDatabase()
	utils.InitDatabase()

	t.Run("should be a product list", func(t *testing.T) {
		// given
		db := utils.GetDB()
		repo := product.NewRepository(db)
		product := domain.Product{
			Name:     "batata",
			Type:     "vegetal",
			Quantity: 1,
			Price:    6.0,
		}

		// when
		err := repo.Save(&product)

		// then

		assert.Nil(t, err)
		assert.NotEqual(t, product.Id, 0)

	})
}
