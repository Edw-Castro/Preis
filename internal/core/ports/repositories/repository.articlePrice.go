package ports

import "github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"

type ArticlePriceRepository interface {
	GetByNameBrand(name, brand string) ([]dto.ArticlePrice, error)
}
