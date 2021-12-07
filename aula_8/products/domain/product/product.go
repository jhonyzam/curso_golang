package product

import (
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
)

type service struct {
	productStorage domain.ProductStorage
}

func NewService(productStorage domain.ProductStorage) *service {
	return &service{
		productStorage: productStorage,
	}
}

func (s *service) Create(product *domain.Product) (*domain.Product, error) {
	return s.productStorage.Insert(product)
}

func (s *service) Get(productID int) (*domain.Product, error) {
	if productID == 0 {
		return nil, domain.ErrInvalidProductID
	}
	return s.productStorage.FindByID(productID)
}

func (s *service) Update(product *domain.Product) error {
	return s.productStorage.Update(product)
}

func (s *service) Delete(productID int) error {
	return s.productStorage.Delete(productID)
}

func (*service) IsValid(product *domain.Product) bool {
	return product.Name != "" && product.Price > 0
}
