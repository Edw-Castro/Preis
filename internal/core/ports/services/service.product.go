package ports

import "github.com/Edw-Castro/Preis/internal/core/domain"

type ProductService interface {
	Detail(id int) (domain.Product, error)
	// GetAllProducts(id ...int) ([]domain.Product, error)
	// GetProduct(id int) (domain.Product, error)
	// CreateProduct(product domain.Product)
	// UpdateProduct(id int) (domain.Product, error)
	// DeleteProduct(id int) error
	// CompareProducts(filter ...string) ([]domain.Product, error)
}
