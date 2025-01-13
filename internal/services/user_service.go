// internal/services/user_service.go
package services

import (
	"errors"
	"finance_backend/internal/models"
	"finance_backend/internal/repository"
	"finance_backend/pkg/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetProfile(userID uint) (*models.User, error) {
	return s.repo.FindByID(userID)
}

func (s *UserService) UpdateProfile(user *models.User) error {
	return s.repo.Update(user)
}
