package service

import (
	"errors"
	"mood-api/internal/models"
	"mood-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("senha inválida")
	}

	return user, nil
}
