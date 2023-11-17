package ports

import "github.com/Edw-Castro/Preis/internal/core/domain"

type BusinessOwnerService interface {
	SignUp(businessOwner *domain.User) error
	// GetAllProducts(id ...int) ([]domain.Product, error)
	// GetProduct(id int) (domain.Product, error)
	// CreateProduct(product domain.Product)
	// UpdateProduct(id int) (domain.Product, error)
	// DeleteProduct(id int) error
	// CompareProducts(filter ...string) ([]domain.Product, error)
}
