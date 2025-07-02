package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type Product struct {
	ID   int
	Name string
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(database:3306)/bgw7")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows, err := db.Query("SELECT product_id, name FROM products")
		if err != nil {
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ID, &p.Name); err != nil {
				http.Error(w, "Error scanning row", http.StatusInternalServerError)
				return
			}
			products = append(products, p)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Error with rows", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}
