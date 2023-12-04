package sql

import (
	"fmt"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	"gorm.io/gorm"
)

type Gorm *gorm.DB

type Repository struct {
	db Gorm
}

func NewProductRepository(db Gorm) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByID(id int) (domain.Product, error) {
	fmt.Println("parametro de GetByID", id)
	product := &domain.Product{}
	// Realizar una consulta SQL para obtener la persona por ID
	repo := *r.db
	result := repo.Table("article").Where("article_id = ?", id).Select("article_id,name_article,brand,category,content,unitMeasurement").Scan(&product)
	fmt.Println("result es:", result)
	fmt.Println("product es:", product)
	return *product, nil
}

func (r *Repository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product

	// Realizar una consulta SQL para obtener todos los productos
	repo := *r.db
	result := repo.Table("article").Select("name_article, brand, category").Scan(&products)

	fmt.Println("result es:", result)
	fmt.Println("products son:", products)

	return products, nil
}
