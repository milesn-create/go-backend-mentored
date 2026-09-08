package main

import "errors"

type UserService struct {
	repo *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{
		repo: r,
	}

}

func (s *UserService) CreateUser(user User) (User, error) {
	u, exists := s.repo.FindByName(user.Name)
	if exists {
		return u, errors.New("user already exists")

	}
	return s.repo.Create(user), nil

}
