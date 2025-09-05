// package product_test

// import (
// 	"regexp"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/bgw7/products-api/internal/domain"
// 	"github.com/bgw7/products-api/internal/product"
// 	"github.com/bgw7/products-api/tests/utils"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_MySqlRepositoryWithSqlMock_GetOne_Mock(t *testing.T) {
// 	productId := 1
// 	db, mock := utils.InitSqlMockDatabase(t)
// 	defer db.Close()

// 	columns := []string{
// 		"id",
// 		"name",
// 		"price",
// 		"quantity",
// 		"type",
// 	}

// 	rows := sqlmock.NewRows(columns)
// 	rows.AddRow(productId, "", 0.0, 0, "")

// 	mock.
// 		ExpectQuery(regexp.QuoteMeta(product.GetQuery)).
// 		WithArgs(productId).
// 		WillReturnRows(rows)

// 	repository := product.NewRepository(db)
// 	result, err := repository.GetByID(productId)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, productId, result.Id)
// }

// func Test_MySqlRepositoryWithSqlMock_Store_Mock(t *testing.T) {
// 	createProduct := domain.Product{
// 		Id: 1,
// 	}

// 	db, mock := utils.InitSqlMockDatabase(t)
// 	defer db.Close()

// 	mock.ExpectExec(regexp.QuoteMeta(product.StoreQuery)).
// 		WithArgs(
// 			createProduct.Name,
// 			createProduct.Price,
// 			createProduct.Quantity,
// 			createProduct.Type,
// 		).WillReturnResult(sqlmock.NewResult(1, 1))

// 	repository := product.NewRepository(db)
// 	repository.Save(&createProduct)
// }

// func Test_MySqlRepositoryWithSqlMock_Delete_Mock(t *testing.T) {
// 	productId := 1
// 	db, mock := utils.InitSqlMockDatabase(t)
// 	defer db.Close()

// 	mock.ExpectExec(regexp.QuoteMeta(product.DeleteQuery)).
// 		WithArgs(productId).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	repository := product.NewRepository(db)
// 	err := repository.Delete(productId)

//		assert.NoError(t, err)
//	}
package product_test
