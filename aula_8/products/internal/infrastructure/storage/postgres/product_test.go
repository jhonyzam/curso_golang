package postgres_test

import (
	"testing"
)

var (
	//testLog = nlog.NewLogger("product-api-test", "", nlog.LevelError)
)

func TestProductStorage_Insert(t *testing.T) {
	t.Run("should insert a new product successfully", func(t *testing.T) {
		//productStorage := postgres.NewProductStorage(TestDb, testLog)
		//assert.NotNil(t, productStorage)
		//
		//var product = &domain.Product{
		//	Name:  "Iso Whey 930g",
		//	Price: 29.90,
		//}
		//
		//product, err := productStorage.Insert(product)
		//assert.NoError(t, err)
		//assert.NotNil(t, product)
		//
		//result := &domain.Product{}
		//err = TestDb.QueryRow(
		//	"SELECT * FROM products WHERE id=$1",
		//	product.ID).Scan(&result.ID, &result.Name, &result.Price)
		//assert.NoError(t, err)
		//assert.Equal(t, product, result)
		//
		//err = productStorage.Delete(product.ID)
		//assert.NoError(t, err)
	})
}

func TestProductStorage_FindByID(t *testing.T) {
	t.Run("should find a product by id", func(t *testing.T) {
		//productStorage := postgres.NewProductStorage(TestDb, testLog)
		//assert.NotNil(t, productStorage)
		//
		//var product = &domain.Product{
		//	Name:  "Iso Whey 930g",
		//	Price: 29.90,
		//}
		//
		//product, err := productStorage.Insert(product)
		//assert.NoError(t, err)
		//assert.NotNil(t, product)
		//
		//searchedProduct, err := productStorage.FindByID(product.ID)
		//assert.NoError(t, err)
		//assert.NotNil(t, searchedProduct)
		//assert.Equal(t, product.Name, searchedProduct.Name)
		//assert.Equal(t, product.Price, searchedProduct.Price)
		//
		//err = productStorage.Delete(product.ID)
		//assert.NoError(t, err)
	})

	t.Run("should not find a product by nonexistent id", func(t *testing.T) {
		//productStorage := postgres.NewProductStorage(TestDb, testLog)
		//assert.NotNil(t, productStorage)
		//
		//var product = &domain.Product{
		//	ID: 9999,
		//}
		//searchedProduct, err := productStorage.FindByID(product.ID)
		//assert.Error(t, err)
		//assert.Nil(t, searchedProduct)
	})
}

func TestProductStorage_Update(t *testing.T) {
	t.Run("should update a product", func(t *testing.T) {
		//productStorage := postgres.NewProductStorage(TestDb, testLog)
		//assert.NotNil(t, productStorage)
		//
		//var product = &domain.Product{
		//	Name:  "Iso Whey 930g",
		//	Price: 29.90,
		//}
		//product, err := productStorage.Insert(product)
		//assert.NoError(t, err)
		//assert.NotNil(t, product)
		//
		//var updatedProduct = &domain.Product{
		//	ID:    product.ID,
		//	Name:  product.Name,
		//	Price: product.Price + 5.10,
		//}
		//
		//err = productStorage.Update(updatedProduct)
		//assert.NoError(t, err)
		//
		//result, err := productStorage.FindByID(updatedProduct.ID)
		//assert.NoError(t, err)
		//assert.NotNil(t, result)
		//assert.Equal(t, updatedProduct.Price, result.Price)
		//assert.NotEqual(t, product.Price, result.Price)
		//
		//err = productStorage.Delete(product.ID)
		//assert.NoError(t, err)
	})
}

func TestProductStorage_Delete(t *testing.T) {
	t.Run("should delete a product ", func(t *testing.T) {
		//productStorage := postgres.NewProductStorage(TestDb, testLog)
		//assert.NotNil(t, productStorage)
		//
		//var product = &domain.Product{
		//	Name:  "Iso Whey 930g",
		//	Price: 29.90,
		//}
		//product, err := productStorage.Insert(product)
		//assert.NoError(t, err)
		//assert.NotNil(t, product)
		//
		//err = productStorage.Delete(product.ID)
		//assert.NoError(t, err)
		//
		//result := -1
		//err = TestDb.QueryRow(
		//	"SELECT count(*) FROM products").Scan(&result)
		//assert.NoError(t, err)
		//assert.Equal(t, 0, result)
	})
}
