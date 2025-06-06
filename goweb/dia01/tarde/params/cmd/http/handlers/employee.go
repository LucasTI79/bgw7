package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ControllerEmployee is the controller for the employee entity that returns handlers
type handlerEmployee struct {
	// storage
	st map[string]string // key: id, value: employee name
}

type ResponseGetByIdEmployee struct {
	Message string `json:"message"`
	Data    *struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Error bool `json:"error"`
}

func NewHandlerEmployee() *handlerEmployee {
	return &handlerEmployee{}
}

func (c *handlerEmployee) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		pathId := chi.URLParam(r, "id")

		queryId := r.URL.Query().Get("id")

		fmt.Println("pathId", pathId)
		fmt.Println("queryId", queryId)

		// process
		// -> get employee
		employee, ok := c.st[pathId]
		if !ok {
			code := http.StatusNotFound
			body := &ResponseGetByIdEmployee{Message: "Employee not found", Data: nil, Error: true}

			w.WriteHeader(code)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &ResponseGetByIdEmployee{Message: "Employee found", Data: &struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{Id: pathId, Name: employee}, Error: false}

		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}
}
