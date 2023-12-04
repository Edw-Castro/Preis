package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/repositories"
	"golang.org/x/crypto/bcrypt"
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

func (bos *businessOwnerService) Login(businessOwner *domain.User) (*domain.User, error) {
	user, err := bos.repo.GetClientUserByEmail(businessOwner.Email)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	// Comparar los hashes
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(businessOwner.Password)); err != nil {
		return nil, err
	}

	return &user, nil
}
