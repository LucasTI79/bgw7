package product_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bgw7/products-api/cmd/http/controller"
	"github.com/bgw7/products-api/config"
	"github.com/bgw7/products-api/internal/product"
	"github.com/bgw7/products-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func Test_GetProduct(t *testing.T) {
	config.Init()
	utils.RegisterDatabase("../../../.env.test")
	utils.InitDatabase()
	defer utils.GetDB().Close()

	t.Run("Test_GetProduct_OK", func(t *testing.T) {
		// given
		db := utils.GetDB()
		repository := product.NewRepository(db)
		controller := controller.NewProductController(repository)
		expectedStatusCode := http.StatusNoContent

		req, res := createRequestTest(t, http.MethodGet, "/products", "")

		// when
		handler := controller.Index()
		handler.ServeHTTP(res, req)

		// then
		assert.Equal(t, expectedStatusCode, res.Code)
	})
}

func createRequestTest(t *testing.T, method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	t.Helper()
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
