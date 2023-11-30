package ports

type ArticlePriceService interface {
	GetByNameBrand(name, brand string) (ArticlePrice, error)
}
