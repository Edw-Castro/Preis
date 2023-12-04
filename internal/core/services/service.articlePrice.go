package services

import (
	ports "github.com/Edw-Castro/Preis/internal/core/ports/repositories"
	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
)

type articelePriceService struct {
	repo ports.ArticlePriceRepository
}

// // GetByNameBrand implements ports.ArticlePriceService.
// func (*articelePriceService) GetByNameBrand(name string, brand string) ([]dto.ArticlePrice, error) {
// 	panic("unimplemented")
// }

// // Detail implements ports.ProductService.
// func (*articelePriceService) Detail(id int) (domain.Product, error) {
// 	panic("unimplemented")
// }

// // GetAllProducts implements ports.ProductService.
// func (*articelePriceService) GetAllProducts() ([]domain.Product, error) {
// 	panic("unimplemented")
// }

func NewArticlePriceService(repo ports.ArticlePriceRepository) *articelePriceService {
	return &articelePriceService{
		repo: repo,
	}
}

func (as *articelePriceService) Compare(name, brand string) ([]dto.ArticlePrice, error) {
	return as.repo.GetByNameBrand(name, brand)
}
