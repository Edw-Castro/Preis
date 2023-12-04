package services

import (
	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/repositories"
)

type productService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *productService {
	return &productService{
		repo: repo,
	}
}

func (ps *productService) Detail(id int) (domain.Product, error) {
	return ps.repo.GetByID(id)
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	return ps.repo.GetAllProducts()
}
