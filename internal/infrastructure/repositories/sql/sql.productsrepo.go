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
	product := &domain.Product{}
	// Realizar una consulta SQL para obtener la persona por ID
	repo := *r.db
	result := repo.Table("article").Where("article_id = ?", id).Select("article_id,name_article,brand,category,content,unitMeasurement").Scan(&product)
	fmt.Println("Lo que encontró la base de datos", result)

	return *product, nil
}
