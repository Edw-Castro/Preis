package ports

import (
	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
)

type ArticlePriceService interface {
	Compare(name, brand string) ([]dto.ArticlePrice, error)
}
