package service

import (
	"mood-api/internal/models"
	"mood-api/internal/repository"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) SignUp(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}
