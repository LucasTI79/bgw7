package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.com/bgw7/http-params/cmd/http/handlers"
)

// Path Param => /:id => utilizado quando queremos buscar um recurso em expecifico
// Query Params => ?name=lucas&age=22&page=1&pageSize=15 => paginacao e filtros

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	empolyeeHandler := handlers.NewHandlerEmployee()

	router.Get("/employees/{id}", empolyeeHandler.GetById())

	log.Fatal(http.ListenAndServe(":8080", router))
}
