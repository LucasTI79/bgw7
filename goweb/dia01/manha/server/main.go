package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// localStorage
// sessionStorage
// cookies

func main() {
	// server
	rt := chi.NewRouter()

	rt.Use(middleware.Heartbeat(""))

	// rt.Route("/api/v1", func(r chi.Router) {
	// 	r.Use(middleware.Logger)
	// })

	// -> endpoints
	rt.
		// With(middleware.Logger).
		Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
			// set code and body
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello World!"))
		})
	// run
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
