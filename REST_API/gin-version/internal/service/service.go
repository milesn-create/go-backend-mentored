package service

import (
	"errors"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/repository"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserService struct {
	repo repository.UserRepository
}
type OrderService struct {
	repo repository.OrderRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}

}
func NewOrderService(r repository.OrderRepository) *OrderService {
	return &OrderService{repo: r}
}

var ErrNameAlreadyExists = errors.New("this name already exists")
var ErrInvalidAge = errors.New("age must be > 0")

func (s *OrderService) CreateOrder(order models.CreateOrderRequest) (models.OrderResponse, error) {
	var pgErr *pgconn.PgError
	resultResponse, err := s.repo.CreateOrder(order)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return models.OrderResponse{}, repository.ErrUserNotFound
		}
		return models.OrderResponse{}, err

	}
	return resultResponse, nil

}
func (s *UserService) UpdateUser(id int, fields models.UserUpdate) (models.User, error) {
	if fields.Age != nil && *fields.Age <= 0 {
		return models.User{}, ErrInvalidAge
	}

	if fields.Name != nil {
		user, err := s.FindByID(id)
		if err != nil {
			return models.User{}, err
		}
		if user.Name != *fields.Name {
			_, exists, err := s.repo.FindByName(*fields.Name)
			if err != nil {
				return models.User{}, err
			}
			if exists {
				return models.User{}, ErrNameAlreadyExists
			}

		}

	}

	updateUser, err := s.repo.UpdateField(id, fields)
	if err != nil {
		return models.User{}, err
	}
	return updateUser, nil

}
func (s *UserService) DeleteUser(id int) (models.User, error) {
	user, err := s.repo.DeleteUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
func (s *UserService) FindByID(id int) (models.User, error) {
	user, exists, err := s.repo.FindByID(id)
	if err != nil {
		return models.User{}, err
	}
	if !exists {
		return models.User{}, repository.ErrUserNotFound
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
