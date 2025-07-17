package router

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/controller"
	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/internal/product"
	"github.com/go-chi/chi/v5"
)

func buildProductRoutes() http.Handler {
	r := chi.NewRouter()

	db := make(map[int]domain.ProductAttributes)
	repository := product.NewRepository(db)
	productController := controller.NewProductController(repository)

	r.Get("/", productController.Index())
	r.Get("/{productId}", productController.Show())
	r.Post("/", productController.Store())
	r.Put("/{productId}", productController.Update())
	r.Patch("/{productId}", productController.Patch())
	r.Delete("/{productId}", productController.Delete())

	return r
}
