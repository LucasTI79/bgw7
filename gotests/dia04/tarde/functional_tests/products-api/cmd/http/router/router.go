package router

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/middlewares"
	"github.com/go-chi/chi/v5"
)

type router struct{}

func NewRouter() *router {
	return &router{}
}

func (router router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.JsonMiddleware)

	r.Route("/products", func(r chi.Router) {
		r.Mount("/", buildProductRoutes())
	})

	return r
}
