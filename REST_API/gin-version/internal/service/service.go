package service

import (
	"errors"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}

}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	u, exists := s.repo.FindByName(user.Name)
	if exists {
		return u, errors.New("user already exists")

	}
	return s.repo.Create(user), nil

}
