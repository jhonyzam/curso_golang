package domain

type Product struct {
	ID    int     `json:"productId,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type ProductService interface {
	Create(product *Product) (*Product, error)
	Get(productID int) (*Product, error)
	Update(product *Product) error
	Delete(productID int) error
	IsValid(product *Product) bool
}

type ProductStorage interface {
	Insert(product *Product) (*Product, error)
	FindByID(productID int) (*Product, error)
	Update(product *Product) error
	Delete(productID int) error
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}
