package main

import (
	"encoding/json"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, _ *http.Request) {
	js, err := json.Marshal(map[string]interface{}{
		"name": "DH-Service",
		"info": "It is ok",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", HandleHealth)
	http.ListenAndServe(":8080", nil)
}
