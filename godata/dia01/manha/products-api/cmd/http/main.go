package main

import (
	"net/http"

	"github.com/bgw7/products-api/cmd/http/router"
	"github.com/bgw7/products-api/config"
	"github.com/bgw7/products-api/database"
)

// Data Transfer Object
// Mapper

func main() {
	config.Init()

	if err := database.Initialize(); err != nil {
		panic(err)
	}

	r := router.NewRouter()

	if err := http.ListenAndServe(":8080", r.MapRoutes()); err != nil {
		panic(err)
	}
}
