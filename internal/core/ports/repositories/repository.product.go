package ports

import "github.com/Edw-Castro/Preis/internal/core/domain"

type ProductRepository interface {
	// Create(product domain.Product) error
	GetByID(id int) (domain.Product, error)
	// getByCategory(category string)
	// Put(id int) (domain.Product, error)
	// Delete(id int) error
}
