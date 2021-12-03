package product_test

import (
	"encoding/json"
	"fmt"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeProduct = &domain.Product{
	ID:    102,
	Name:  "Iso Whey 930g",
	Price: 29.90,
}

var fakeCompany = &domain.Company{
	CNPJ: "78079606000151",
}

func TestService_Create(t *testing.T) {
	productInsert := *fakeProduct

	t.Run("should create a product", func(t *testing.T) {

		productStorageMock := &product.StorageMock{
			InsertFn: func(product *domain.Product) (*domain.Product, error) {
				assert.Equal(t, &productInsert, product)
				return product, nil
			},
		}

		productService := product.NewService(productStorageMock)
		product, err := productService.Create(&productInsert)

		assert.Equal(t, 1, productStorageMock.InsertInvokedCount)
		assert.NoError(t, err)
		assert.Equal(t, &productInsert, product)
	})

	t.Run("should return error from storage", func(t *testing.T) {

		productStorageMock := &product.StorageMock{
			InsertFn: func(product *domain.Product) (*domain.Product, error) {
				return product, fmt.Errorf("mock error")
			},
		}

		productService := product.NewService(productStorageMock)
		product, err := productService.Create(&productInsert)

		assert.Equal(t, 1, productStorageMock.InsertInvokedCount)
		assert.Errorf(t, err, "mock error")
		assert.Equal(t, &productInsert, product)
	})
}

func TestService_Get(t *testing.T) {
	productFind := *fakeProduct

	t.Run("should get a product", func(t *testing.T) {

		productStorageMock := &product.StorageMock{
			FindByIDFn: func(productID int) (*domain.Product, error) {
				return &productFind, nil
			},
		}

		productService := product.NewService(productStorageMock)
		assert.NotNil(t, productService)

		product, err := productService.Get(102)
		assert.Equal(t, 1, productStorageMock.FindByIDInvokedCount)
		assert.NoError(t, err)
		assert.Equal(t, &productFind, product)
	})

	t.Run("should return productId invalid error", func(t *testing.T) {

		productStorageMock := &product.StorageMock{
			FindByIDFn: func(productID int) (*domain.Product, error) {
				return nil, nil
			},
		}

		productService := product.NewService(productStorageMock)
		assert.NotNil(t, productService)

		product, err := productService.Get(0)
		assert.Nil(t, product)
		assert.Error(t, err)
		assert.Equal(t, "productId is invalid", err.Error())
		assert.Equal(t, 0, productStorageMock.FindByIDInvokedCount)

	})

}

func TestService_Update(t *testing.T) {
	productUpdate := *fakeProduct

	t.Run("should update a product", func(t *testing.T) {
		productStorageMock := &product.StorageMock{
			UpdateFn: func(product *domain.Product) error {
				assert.Equal(t, &productUpdate, product)
				return nil
			},
		}

		productService := product.NewService(productStorageMock)
		assert.NotNil(t, productService)

		err := productService.Update(&productUpdate)
		assert.NoError(t, err)
		assert.Equal(t, 1, productStorageMock.UpdateInvokedCount)
	})

	t.Run("should return error from storage", func(t *testing.T) {
		productStorageMock := &product.StorageMock{
			UpdateFn: func(product *domain.Product) error {
				return fmt.Errorf("mock error")
			},
		}

		productService := product.NewService(productStorageMock)
		assert.NotNil(t, productService)

		err := productService.Update(&productUpdate)
		assert.Equal(t, 1, productStorageMock.UpdateInvokedCount)
		assert.Errorf(t, err, "mock error")
	})

}

func TestService_Delete(t *testing.T) {

	t.Run("should delete a product", func(t *testing.T) {
		productStorageMock := &product.StorageMock{
			DeleteFn: func(productID int) error {
				return nil
			},
		}

		productService := product.NewService(productStorageMock)
		assert.NotNil(t, productService)

		err := productService.Delete(1)
		assert.NoError(t, err)
		assert.Equal(t, 1, productStorageMock.DeleteInvokedCount)
	})

	t.Run("should return productId invalid error", func(t *testing.T) {

		productStorageMock := &product.StorageMock{
			FindByIDFn: func(productID int) (*domain.Product, error) {
				return nil, nil
			},
		}

		productService := product.NewService(productStorageMock)
		assert.NotNil(t, productService)

		product, err := productService.Get(0)
		assert.Nil(t, product)
		assert.Error(t, err)
		assert.Equal(t, "productId is invalid", err.Error())
		assert.Equal(t, 0, productStorageMock.FindByIDInvokedCount)

	})

}

func TestService_IsValid(t *testing.T) {

	t.Run("product must be valid", func(t *testing.T) {
		validProduct := domain.NewProduct("Budweiser Beer 355ml", 4.15)

		productService := product.NewService(nil)
		assert.NotNil(t, productService)
		assert.True(t, productService.IsValid(validProduct))
	})

	t.Run("product must be invalid", func(t *testing.T) {
		validProduct := domain.NewProduct("Budweiser Beer 355ml", 0)

		productService := product.NewService(nil)
		assert.NotNil(t, productService)
		assert.False(t, productService.IsValid(validProduct))
	})

}

func TestProductJson(t *testing.T) {
	product := domain.NewProduct("Budweiser Beer 355ml", 4.15)

	productBytes, err := json.Marshal(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, productBytes)

	expected := "{\"name\":\"Budweiser Beer 355ml\", \"price\":4.15}"
	actual := string(productBytes)

	assert.JSONEq(t, expected, actual)
}
