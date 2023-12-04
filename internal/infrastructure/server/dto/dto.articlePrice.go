package dto

type ArticlePrice struct {
	Name_store   string  `json:"name_store"`
	Name_article string  `json:"name_article"`
	Price        float32 `json:"price"`
}

type ResponseArticlePrice struct {
	ArticleCompare []ArticlePrice
}

func BuildResponseArticlePrice(articlesPrice []ArticlePrice) ResponseArticlePrice {
	response := ResponseArticlePrice{
		ArticleCompare: articlesPrice,
	}
	return response
}
