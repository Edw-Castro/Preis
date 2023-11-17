package ports

import "github.com/Edw-Castro/Preis/internal/core/domain"

type BusinessOwnerRepository interface {
	// Create(product domain.Product) error
	Insert(businessOwner *domain.User) error
	GetClientUserByEmail(email string) (domain.User, error)
}
