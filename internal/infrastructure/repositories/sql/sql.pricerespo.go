package sql

import (
	"fmt"

	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
)

type RepositoryArticle struct {
	db1 Gorm
	db2 Gorm
}

func NewArticleRepository(db1, db2 Gorm) *RepositoryArticle {
	return &RepositoryArticle{
		db1: db1,
		db2: db2,
	}
}
func (r *RepositoryArticle) GetByNameBrand(name, brand string) ([]dto.ArticlePrice, error) {
	var articles []dto.ArticlePrice

	database1 := *r.db1
	database2 := *r.db2

	// Realiza la consulta en ambas bases de datos
	result := database1.Table("store_article_price p").
		Joins("JOIN store s ON p.store_id = s.store_id").
		Joins("JOIN ?", database2.
			Table("article a").
			Where("a.name_article = ? AND a.brand = ?", name, brand).
			Select("a.article_id")).
		Where("p.article_id = a.article_id").
		Select("s.name_store, a.name_article, p.price").
		Scan(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Se encontraron", len(articles), "registros")

	return articles, nil
}

// func (r *RepositoryArticle) GetByNameBrand(name, brand string) ([]dto.ArticlePrice, error) {
// 	var articles []dto.ArticlePrice
// 	/*database1 := *r.db2
// 	fmt.Println("REEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE GETBYNAMEBRAND")
// 	result := database1.Table("store_article_price").
// 		Joins("JOIN article ON store_article_price.article_id = article.id").
// 		Joins("JOIN store ON store_article_price.store_id = store.id").
// 		Where("product.name_article = ?", name).
// 		Select("store.name_store, product.name, store_article_price.price,store_article_price.price/store_article_price.content").
// 		Order("store_article_price.price/store_article_price.content DESC").
// 		Limit(10).Scan(&articles)
// 	fmt.Println("Se encontro", result)*/

// 	database1 := *r.db1
// 	var articleIDs []string
// 	result1 := database1.Table("article").
// 		Where("name_article = ?", name).
// 		Scan(&articleIDs)
// 	println("Los nombres que encontro son: ", articleIDs)
// 	if result1.Error != nil {
// 		return nil, result1.Error
// 	}

// 	// Realiza la segunda consulta en la base de datos 2 (r.db2)
// 	database2 := *r.db2
// 	result2 := database2.Table("store_article_price").
// 		Joins("JOIN store ON store_article_price.store_id = store.id").
// 		Where("article_id IN ?", articleIDs).
// 		Select("store.name_store, article.name_article, store_article_price.price, store_article_price.price/store_article_price.content").
// 		Order("store_article_price.price/store_article_price.content DESC").
// 		Limit(10).Scan(&articles)

// 	if result2.Error != nil {
// 		return nil, result2.Error
// 	}

// 	fmt.Println("Se encontraron", len(articles), "registros")

// 	return articles, nil
// }
