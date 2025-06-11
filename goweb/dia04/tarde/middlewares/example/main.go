package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func CheckTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before
		fmt.Println("Request received at", time.Now())

		// call handler
		handler.ServeHTTP(w, r)

		// do something after
	})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

func Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before
		token := r.Header.Get("token")
		if token != "secret" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Unauthorized"))
			return
		}
		// call handler
		handler.ServeHTTP(w, r)
		// do something after
	})
}

func main() {
	// server
	rt := chi.NewRouter()

	// middleware global
	rt.Use(CheckTime)
	rt.Get("/health", HealthCheck)

	rt.
		With(Auth).
		Get("/", HelloWorld)

	http.ListenAndServe(":8080", rt)
}
