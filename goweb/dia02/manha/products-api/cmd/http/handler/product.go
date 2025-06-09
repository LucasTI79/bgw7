package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bgw7/products-api/internal/domain"
)

// ControllerProducts is a struct that contains the storage of products
type ProductHandler struct {
	storage map[int]*domain.Product
}

func NewProductHandler(db map[int]*domain.Product) *ProductHandler {
	return &ProductHandler{storage: db}
}

func (c *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &domain.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// request
		var reqBody domain.RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			code := http.StatusBadRequest
			body := &domain.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = len(c.storage) + 1

		// pr := &domain.Product{
		// 	Id:       len(c.storage) + 1,
		// 	Name:     reqBody.Name,
		// 	Type:     reqBody.Type,
		// 	Quantity: reqBody.Quantity,
		// 	Price:    reqBody.Price,
		// }
		// -> save product
		c.storage[pr.Id] = &pr

		// response
		code := http.StatusCreated
		body := &domain.ResponseBodyProduct{
			Message: "Product created",
			Data: &domain.Product{
				Id:       pr.Id,
				Name:     pr.Name,
				Type:     pr.Type,
				Quantity: pr.Quantity,
				Price:    pr.Price,
			},
			Error: false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}
