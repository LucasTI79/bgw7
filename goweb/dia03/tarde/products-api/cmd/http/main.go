package main

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/controller"
	"github.com/bgw7/products-api/config"
	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/internal/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Data Transfer Object
// Mapper

func main() {
	config.Init()
	rt := chi.NewRouter()

	rt.Use(middleware.Logger)

	rt.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	db := make(map[int]domain.ProductAttributes)
	repository := product.NewRepository(db)
	productController := controller.NewProductController(repository)

	// -> rotas
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", productController.Index())
		r.Get("/{productId}", productController.Show())
		r.Post("/", productController.Store())
		r.Put("/{productId}", productController.Update())
		r.Patch("/{productId}", productController.Patch())
		r.Delete("/{productId}", productController.Delete())
	})

	// -> executar
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
