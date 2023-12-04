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

func (r *Repository) Insert(user *domain.User) error {

	repo := *r.db
	result := repo.Table("user").Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) GetClientUserByEmail(email string) (domain.User, error) {
	fmt.Println("El mail es", email)
	client := &domain.User{}
	repo := *r.db
	repo.Table("user").Where("mail = ?", email).Select("user_id, name_user, mail, rol, pwd").Scan(&client)
	fmt.Println("El cliente o usaurio obtenido es:", client)
	return *client, nil
}
