package service

import (
	"errors"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}

}

var ErrUserNotFound = errors.New("user not found\n")

func (s *UserService) FindByID(id int) (models.User, error) {
	user, exists, err := s.repo.FindByID(id)
	if err != nil {
		return models.User{}, err
	}
	if !exists {
		return models.User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	u, exists, err := s.repo.FindByName(user.Name)
	if err != nil {
		return models.User{}, err
	}
	if exists {
		return u, errors.New("user already exists")

	}
	return s.repo.Create(user)

}
