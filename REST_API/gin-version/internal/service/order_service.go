package service

import (
	"errors"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/repository"

	"github.com/jackc/pgx/v5/pgconn"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) *OrderService {
	return &OrderService{repo: r}
}
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
