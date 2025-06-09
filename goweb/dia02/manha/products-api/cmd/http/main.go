package main

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/handler"
	"github.com/bgw7/products-api/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Data Transfer Object
// Mapper

func main() {
	rt := chi.NewRouter()

	rt.Use(middleware.Logger)

	rt.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	db := make(map[int]*domain.Product)
	productHandler := handler.NewProductHandler(db)

	// -> rotas
	rt.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.Create())
	})

	// -> executar
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
