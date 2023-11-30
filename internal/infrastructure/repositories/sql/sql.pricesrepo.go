package sql

import (
	"fmt"
)

type RepositoryArticle struct {
	db1 Gorm
	db2 Gorm
}
type articlePrice struct {
	name_store   string
	name_article string
	price        float32
}

func NewArticleRepository(db1, db2 Gorm) *RepositoryArticle {
	return &RepositoryArticle{
		db1: db1,
		db2: db2,
	}
}

func (r *RepositoryArticle) GetByNameBrand(name, brand string) ([]articlePrice, error) {
	var articles []articlePrice
	database1 := *r.db1
	result := database1.Table("storeArticlePrice").
		Joins("JOIN article ON store_article_price.article_id = article.id").
		Joins("JOIN store ON store_article_price.store_id = store.id").
		Where("product.name_article = ?", name).
		Select("store.name_store, product.name, store_article_price.price,store_article_price.price/store_article_price.content").
		Order("store_article_price.price/store_article_price.content DESC").
		Limit(10).Scan(&articles)
	fmt.Println("Se encontro", result)
	return articles, nil
}
