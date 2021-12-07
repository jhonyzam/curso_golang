package http_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain/product"
	internalHttp "github.com/jhonyzam/curso_golang/aula_8/products/internal/infrastructure/server/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	fmt.Println("init...")
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestController(t *testing.T) {
	t.Run("should success to post new product", func(t *testing.T) {
		productServiceMock := &product.ServiceMock{
			IsValidFn: func(product *domain.Product) bool {
				return product.Name != "" && product.Price > 0
			},
			CreateFn: func(product *domain.Product) (*domain.Product, error) {
				product.ID = 102
				return product, nil
			},
		}

		router := internalHttp.NewHandler(productServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/products"
		body := []byte(`{"name": "Iso Whey 930g", "price": 29.90}`)

		req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)

		assert.Equal(t, http.StatusCreated, response.Code)

		body, err := ioutil.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		expectedBody := []byte(`{"productId": 102, "name": "Iso Whey 930g", "price": 29.90}`)
		assert.JSONEq(t, string(expectedBody), string(body))
		assert.Equal(t, 1, productServiceMock.IsValidInvokedCount)
	})

	t.Run("should fail to post invalid new product", func(t *testing.T) {
		productServiceMock := &product.ServiceMock{
			IsValidFn: func(product *domain.Product) bool {
				return product.Name != "" && product.Price > 0
			},
		}

		router := internalHttp.NewHandler(productServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/products"
		body := []byte(`{"name": "Iso Whey 930g", "price": -10.0}`)

		req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, 1, productServiceMock.IsValidInvokedCount)
	})

	t.Run("should success to get existent product by id", func(t *testing.T) {
		productServiceMock := &product.ServiceMock{
			GetFn: func(productID int) (*domain.Product, error) {
				product := &domain.Product{
					ID:    1,
					Name:  "Iso Whey 930g",
					Price: 29.90,
				}
				if productID == product.ID {
					return product, nil
				}
				return nil, errors.New("Couldn't find product")
			},
		}

		router := internalHttp.NewHandler(productServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/products/1"
		body := []byte(`{}`)

		req, _ := http.NewRequest("GET", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, 1, productServiceMock.GetInvokedCount)
	})

	t.Run("should fail to get nonexistent product by id", func(t *testing.T) {
		productServiceMock := &product.ServiceMock{
			GetFn: func(productID int) (*domain.Product, error) {
				product := &domain.Product{
					ID:    1,
					Name:  "Iso Whey 930g",
					Price: 29.90,
				}
				if productID == product.ID {
					return product, nil
				}
				return nil, errors.New("Couldn't find product")
			},
		}

		router := internalHttp.NewHandler(productServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/products/99999999"
		body := []byte(`{}`)

		req, _ := http.NewRequest("GET", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, 1, productServiceMock.GetInvokedCount)
	})

	t.Run("should success to update product by id", func(t *testing.T) {
		productServiceMock := &product.ServiceMock{
			IsValidFn: func(product *domain.Product) bool {
				return product.Name != "" && product.Price > 0
			},
			UpdateFn: func(product *domain.Product) error {
				return nil
			},
		}

		router := internalHttp.NewHandler(productServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/products/1"
		body := []byte(`{"name": "Iso Whey 930g", "price": 39.90}`)

		req, _ := http.NewRequest("PUT", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusNoContent, response.Code)
		assert.Equal(t, 1, productServiceMock.UpdateInvokedCount)
	})

	t.Run("should success to delete product by id", func(t *testing.T) {
		productServiceMock := &product.ServiceMock{
			DeleteFn: func(productID int) error {
				return nil
			},
		}

		router := internalHttp.NewHandler(productServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/products/1"
		body := []byte(`{}`)

		req, _ := http.NewRequest("DELETE", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusNoContent, response.Code)
		assert.Equal(t, 1, productServiceMock.DeleteInvokedCount)
	})
}
