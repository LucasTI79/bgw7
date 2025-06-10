package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bgw7/products-api/config"
	"github.com/bgw7/products-api/internal/domain"
	"github.com/bgw7/products-api/internal/plataform/input"
	"github.com/bgw7/products-api/internal/plataform/output"
	"github.com/bgw7/products-api/pkg/apperrors"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// ControllerProducts is a struct that contains the storage of products
type ProductController struct {
	repository domain.ProductRepository
}

func NewProductController(repository domain.ProductRepository) *ProductController {
	return &ProductController{repository: repository}
}

func (c *ProductController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetEnv[string]("ApiToken")

		fmt.Println(apiToken)

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := &output.CreateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		products, err := c.repository.Get()

		if err != nil {
			code := http.StatusInternalServerError
			body := map[string]any{"message": "internal server error", "data": nil}
			response.JSON(w, code, body)
			return
		}

		if len(products) == 0 {
			code := http.StatusNoContent
			response.JSON(w, code, nil)
			return
		}

		// response
		code := http.StatusOK
		body := &output.GetBodyProductOutput{
			Message: "products list",
			Data:    products,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

func (c *ProductController) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetConfig().ApiToken

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := &output.CreateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			response.JSON(w, code, body)
			return
		}

		product, err := c.repository.GetByID(productId)

		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": err.Error(), "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &output.GetOneBodyProductOutput{
			Message: "product",
			Data:    product,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

func (c *ProductController) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetConfig().ApiToken

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := &output.CreateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// request
		var reqBody input.CreateBodyProductInput
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := &output.CreateBodyProductOutput{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()

		// pr := &domain.Product{
		// 	Id:       len(c.storage) + 1,
		// 	Name:     reqBody.Name,
		// 	Type:     reqBody.Type,
		// 	Quantity: reqBody.Quantity,
		// 	Price:    reqBody.Price,
		// }
		// -> save product
		if err := c.repository.Save(&pr); err != nil {
			code := http.StatusInternalServerError
			body := map[string]any{"message": "internal server error", "data": nil}
			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusCreated
		body := &output.CreateBodyProductOutput{
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

func (c *ProductController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetConfig().ApiToken

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := &output.CreateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			response.JSON(w, code, body)
			return
		}

		var reqBody input.UpdateBodyProductInput
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusUnprocessableEntity
			body := map[string]any{"message": "invalid request body", "data": nil}

			response.JSON(w, code, body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = productId

		if err := c.repository.Update(&pr); err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": err.Error(), "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := map[string]any{"message": "product updated", "data": pr}

		response.JSON(w, code, body)
	}
}

func (c *ProductController) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetConfig().ApiToken

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := &output.CreateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			response.JSON(w, code, body)
			return
		}

		product, err := c.repository.GetByID(productId)
		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		reqBody := input.UpdateBodyProductInput{
			Name:     product.Name,
			Type:     product.Type,
			Quantity: product.Quantity,
			Price:    product.Price,
		}
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusUnprocessableEntity
			body := map[string]any{"message": "invalid request body", "data": nil}

			response.JSON(w, code, body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = productId

		if err := c.repository.Update(&pr); err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": err.Error(), "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := map[string]any{"message": "product updated", "data": pr}

		response.JSON(w, code, body)
	}
}

func (c *ProductController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetConfig().ApiToken

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := &output.CreateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			response.JSON(w, code, body)
			return
		}

		if err := c.repository.Delete(productId); err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": err.Error(), "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		// 404 - Not Found
		// 204 - No Content
		code := http.StatusNoContent
		response.JSON(w, code, nil)
	}
}
