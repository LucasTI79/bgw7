package router

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/controller"
	"github.com/bgw7/products-api/database"
	"github.com/bgw7/products-api/internal/product"
	"github.com/go-chi/chi/v5"
)

func buildProductRoutes() http.Handler {
	r := chi.NewRouter()

	db := database.GetDB()
	repository := product.NewRepository(db)
	productController := controller.NewProductController(repository)

	// /products/info
	// /products/1

	r.Get("/", productController.Index())
	r.Get("/info", productController.IndexInfo())
	r.Get("/{productId}", productController.Show())
	r.Post("/", productController.Store())
	r.Put("/{productId}", productController.Update())
	r.Patch("/{productId}", productController.Patch())
	r.Delete("/{productId}", productController.Delete())

	return r
}
