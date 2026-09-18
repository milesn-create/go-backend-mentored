package repository

import (
	"database/sql"
	"fmt"
	"rest-api-gin/internal/models"
	"strings"
	"time"
)

type OrderRepository interface {
	CreateOrder(order models.CreateOrderRequest) (models.OrderResponse, error)
}
type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(DB *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: DB}
}

func (r *PostgresOrderRepository) CreateOrder(order models.CreateOrderRequest) (models.OrderResponse, error) {
	tr, err := r.db.Begin()
	if err != nil {
		return models.OrderResponse{}, fmt.Errorf("failed create transaction: %w", err)
	}
	defer tr.Rollback()
	columns := []string{"user_id"}
	placeholders := []string{"$1"}
	args := []any{order.UserID}
	if order.Status != nil {
		columns = append(columns, "status")
		args = append(args, *order.Status)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))

	}
	query := "INSERT INTO orders(" + strings.Join(columns, ", ") + ") VALUES(" + strings.Join(placeholders, ", ") + ") RETURNING id, CURRENT_TIMESTAMP"
	var createdAt time.Time
	var orderId int
	err = tr.QueryRow(query, args...).Scan(&orderId, &createdAt)
	if err != nil {
		return models.OrderResponse{}, fmt.Errorf("failed insert into orders: %w", err)
	}
	OrderItems := []models.OrderItem{}
	for i := range order.OrderItems {
		var orderItemId int

		err := tr.QueryRow("INSERT INTO order_items(order_id,product_name,price,quantity) VALUES ($1,$2,$3,$4)RETURNING id", orderId, order.OrderItems[i].ProductName, *order.OrderItems[i].Price, order.OrderItems[i].Quantity).Scan(&orderItemId)
		if err != nil {
			return models.OrderResponse{}, fmt.Errorf("failed insert into order_items: %w", err)
		}
		OrderItems = append(OrderItems, models.OrderItem{ID: orderItemId, OrderId: orderId, ProductName: order.OrderItems[i].ProductName, Price: order.OrderItems[i].Price, Quantity: order.OrderItems[i].Quantity})

	}
	err = tr.Commit()
	if err != nil {
		return models.OrderResponse{}, fmt.Errorf("failed commit transaction : %w", err)
	}

	return models.OrderResponse{Order: models.Order{ID: orderId, UserID: order.UserID, Status: order.Status, CreatedAt: createdAt}, Items: OrderItems}, nil

}
