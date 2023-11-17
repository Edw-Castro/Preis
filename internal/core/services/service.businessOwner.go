package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/repositories"
)

type businessOwnerService struct {
	repo ports.BusinessOwnerRepository
}

func NewBusinessOwnerService(repo ports.BusinessOwnerRepository) *businessOwnerService {
	return &businessOwnerService{
		repo: repo,
	}
}

func (bos *businessOwnerService) SignUp(businessOwner *domain.User) error {
	user, err := bos.repo.GetClientUserByEmail(businessOwner.Email)
	fmt.Println("pasé de user en SignUp")
	if err != nil {
		log.Fatalf("El error es %v", err)
		return err
	}

	verifyEmail := user.Email
	if verifyEmail != "" {
		return errors.New("Users already exists")
	}

	err = bos.repo.Insert(businessOwner)
	if err != nil {
		return err
	}
	fmt.Println("ya se registró")
	return nil
}
