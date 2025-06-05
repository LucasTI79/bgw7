package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func fooHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func main() {
	http.Handle("/foo", fooHandler())

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		// io.Copy(w, strings.NewReader("texto qualquer"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
