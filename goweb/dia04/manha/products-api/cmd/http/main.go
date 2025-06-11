package main

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/router"
	"github.com/bgw7/products-api/config"
)

// Data Transfer Object
// Mapper

func main() {
	config.Init()

	r := router.NewRouter()

	if err := http.ListenAndServe(":8080", r.MapRoutes()); err != nil {
		panic(err)
	}
}
