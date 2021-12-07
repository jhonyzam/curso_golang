package postgres

import (
	"database/sql"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain"
	"log"
)

type productStorage struct {
	DB *sql.DB
}

func NewProductStorage(db *sql.DB) *productStorage {
	return &productStorage{
		DB: db,
	}
}

func (ps *productStorage) Insert(product *domain.Product) (*domain.Product, error) {
	err := ps.DB.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		product.Name, product.Price).Scan(&product.ID)

	if err != nil {
		log.Printf("ERROR: inserting new product: %q\n", err)
		return nil, err
	}
	return product, nil
}

func (ps *productStorage) FindByID(productID int) (*domain.Product, error) {
	var product domain.Product

	err := ps.DB.QueryRow(
		"SELECT id, name, price FROM products WHERE id=$1",
		productID).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		log.Printf("ERROR: finding product: %q\n", err)
		return nil, err
	}
	return &product, nil
}

func (ps *productStorage) Update(product *domain.Product) error {
	_, err := ps.DB.Exec(
		"UPDATE products SET name=$1, price=$2 WHERE id=$3",
		product.Name, product.Price, product.ID)

	if err != nil {
		log.Printf("ERROR: updating product: %q\n", err)
	}

	return err
}

func (ps *productStorage) Delete(productID int) error {
	_, err := ps.DB.Exec(
		"DELETE FROM products WHERE id=$1",
		productID)

	if err != nil {
		log.Printf("ERROR: deleting product: %q", err)
	}

	return err
}
