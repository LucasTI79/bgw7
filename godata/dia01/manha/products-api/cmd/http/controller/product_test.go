package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bgw7/products-api/cmd/http/controller"
	"github.com/bgw7/products-api/cmd/http/middlewares"
	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/internal/product"
	"github.com/bgw7/products-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// docker
// testcontainer

// migrations
// GORM
// SQLC

func TestGetProduct(t *testing.T) {
	url := "/products"

	// F.I.R.S.T

	t.Run("success to get products with statusCode 200", func(t *testing.T) {
		// given

		// test double (fake) => dummies, stubs, spies, mock e fakes
		db := map[int]domain.ProductAttributes{
			1: {Name: "product 1", Type: "type 1", Quantity: 1, Price: 1.1},
			2: {Name: "product 2", Type: "type 2", Quantity: 2, Price: 2.2},
		}
		repository := product.NewRepository(db)
		controller := controller.NewProductController(repository)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusOK
		expectedBody := `{"message":"products list","data":[{"Id":1,"Name":"product 1","Type":"type 1","Quantity":1,"Price":1.1},{"Id":2,"Name":"product 2","Type":"type 2","Quantity":2,"Price":2.2}],"error":false}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// when
		handler := middlewares.JsonMiddleware(controller.Index())
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.JSONEq(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})

	t.Run("success to get products with statusCode 204", func(t *testing.T) {
		db := map[int]domain.ProductAttributes{}
		repository := product.NewRepository(db)
		controller := controller.NewProductController(repository)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusNoContent
		expectedBody := ""
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// when
		handler := middlewares.JsonMiddleware(controller.Index())
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})

	t.Run("failed to get products", func(t *testing.T) {
		repository := mocks.ProductRepositoryMock{}
		controller := controller.NewProductController(&repository)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusInternalServerError
		expectedBody := `{"data":null,"message":"internal server error"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		var mockedProducts []domain.Product
		mockedError := errors.New("failed to connect database")

		// when
		repository.
			On("Get").
			Return(mockedProducts, mockedError)

		handler := middlewares.JsonMiddleware(controller.Index())
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})
}

func TestSaveProduct(t *testing.T) {
	url := "/products"

	// F.I.R.S.T

	t.Run("success to save product with statusCode 201", func(t *testing.T) {
		// given

		// test double (fake) => dummies, stubs, spies, mock e fakes
		db := map[int]domain.ProductAttributes{}
		repository := product.NewRepository(db)
		controller := controller.NewProductController(repository)
		requestBody := `{"name":"LCD","type":"Entertainment","price":20000,"quantity":5}`

		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(requestBody))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusCreated
		expectedBody := `{"message":"product created","data":{"Id":1,"Name":"LCD","Price":20000,"Quantity":5,"Type":"Entertainment"},"error":false}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// when
		handler := middlewares.JsonMiddleware(controller.Store())
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.JSONEq(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})

	t.Run("success to get products with statusCode 400", func(t *testing.T) {
		// given

		// test double (fake) => dummies, stubs, spies, mock e fakes
		db := map[int]domain.ProductAttributes{}
		repository := product.NewRepository(db)
		controller := controller.NewProductController(repository)
		requestBody := ``

		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(requestBody))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusBadRequest
		expectedBody := `{"message":"bad request","data":null,"error":true}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// when
		handler := middlewares.JsonMiddleware(controller.Store())
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.JSONEq(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})

	t.Run("failed to get products with status 500", func(t *testing.T) {
		repository := mocks.ProductRepositoryMock{}
		controller := controller.NewProductController(&repository)
		requestBody := `{"name":"LCD","type":"Entertainment","price":20000,"quantity":5}`

		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(requestBody))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		repository.
			On("Save", mock.Anything).
			Return(errors.New("internal server error"))

		// when
		handler := middlewares.JsonMiddleware(controller.Store())
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
	})
}
