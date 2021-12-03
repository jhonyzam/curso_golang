package product

import (
	"context"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
)

type ServiceMock struct {
	CreateInvokedCount      int
	GetInvokedCount         int
	UpdateInvokedCount      int
	DeleteInvokedCount      int
	IsValidInvokedCount     int
	IsValidCNPJInvokedCount int

	CreateFn      func(product *domain.Product) (*domain.Product, error)
	GetFn         func(productID int) (*domain.Product, error)
	UpdateFn      func(product *domain.Product) error
	DeleteFn      func(productID int) error
	IsValidFn     func(product *domain.Product) bool
	IsValidCNPJFn func(ctx context.Context, CNPJ string) bool
}

func (sm *ServiceMock) Create(product *domain.Product) (*domain.Product, error) {
	sm.CreateInvokedCount++
	return sm.CreateFn(product)
}

func (sm *ServiceMock) Get(productID int) (*domain.Product, error) {
	sm.GetInvokedCount++
	return sm.GetFn(productID)
}

func (sm *ServiceMock) Update(product *domain.Product) error {
	sm.UpdateInvokedCount++
	return sm.UpdateFn(product)
}

func (sm *ServiceMock) Delete(productID int) error {
	sm.DeleteInvokedCount++
	return sm.DeleteFn(productID)
}

func (sm *ServiceMock) IsValid(product *domain.Product) bool {
	sm.IsValidInvokedCount++
	return sm.IsValidFn(product)
}

func (sm *ServiceMock) IsValidCNPJ(ctx context.Context, CNPJ string) bool {
	sm.IsValidCNPJInvokedCount++
	return sm.IsValidCNPJFn(ctx, CNPJ)
}

type StorageMock struct {
	InsertInvokedCount   int
	FindByIDInvokedCount int
	UpdateInvokedCount   int
	DeleteInvokedCount   int

	InsertFn   func(product *domain.Product) (*domain.Product, error)
	FindByIDFn func(productID int) (*domain.Product, error)
	UpdateFn   func(product *domain.Product) error
	DeleteFn   func(productID int) error
}

func (sm *StorageMock) Insert(product *domain.Product) (*domain.Product, error) {
	sm.InsertInvokedCount++
	return sm.InsertFn(product)
}

func (sm *StorageMock) FindByID(productID int) (*domain.Product, error) {
	sm.FindByIDInvokedCount++
	return sm.FindByIDFn(productID)
}

func (sm *StorageMock) Update(product *domain.Product) error {
	sm.UpdateInvokedCount++
	return sm.UpdateFn(product)
}

func (sm *StorageMock) Delete(productID int) error {
	sm.DeleteInvokedCount++
	return sm.DeleteFn(productID)
}
