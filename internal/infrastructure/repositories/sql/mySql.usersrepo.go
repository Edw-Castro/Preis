package sql

import (
	"fmt"

	"github.com/Edw-Castro/Preis/internal/core/domain"
)

func NewUserRepository(db Gorm) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(businessOwner *domain.User) error {

	repo := *r.db
	result := repo.Table("users").Create(&businessOwner)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) GetClientUserByEmail(email string) (domain.User, error) {
	client := &domain.User{}
	// Realizar una consulta SQL para obtener la persona por ID
	repo := *r.db
	result := repo.Table("users").Where("email = ?", email).Select("id, name, email, role").Scan(&client)
	fmt.Println("Lo que encontró la base de datos", result)

	return *client, nil
}
