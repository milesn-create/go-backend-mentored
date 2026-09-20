package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"rest-api-gin/internal/models"
	"strings"
	"time"
)

type OrderRepository interface {
	CreateOrder(order models.CreateOrderRequest) (models.OrderResponse, error)
	GetOrderByID(id int) (models.OrderResponse, error)
}
type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(DB *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: DB}
}

var ErrOrderNotExists = errors.New("Order not exists")

func (r *PostgresOrderRepository) GetOrderByID(id int) (models.OrderResponse, error) {
	rows, err := r.db.Query(
		`SELECT 
			o.id ,
			o.user_id,
    		o.status,
    		o.created_at,
    		oi.id,
    		oi.product_name,
    		oi.price,
    		oi.quantity
		FROM orders o
    		LEFT JOIN order_items oi ON o.id = oi.order_id
		WHERE o.id = $1`, id)
	if err != nil {
		return models.OrderResponse{}, fmt.Errorf("failed SELECT FROM orders LEFT JOIN order_items: %w", err)
	}
	defer rows.Close()
	var order models.Order
	var orderItems []models.OrderItem
	orderExists := false
	for rows.Next() {
		orderExists = true
		var itemID *int
		var productName *string
		var price *int
		var quantity *int

		err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt, &itemID, &productName, &price, &quantity)
		if err != nil {
			return models.OrderResponse{}, fmt.Errorf("failed scan order row: %w", err)
		}
		if itemID != nil {
			orderItems = append(orderItems, models.OrderItem{ID: *itemID, OrderId: order.ID, ProductName: *productName, Price: price, Quantity: *quantity})
		}

	}
	if err := rows.Err(); err != nil {
		return models.OrderResponse{}, fmt.Errorf("failed rows.Next(): %w", err)
	}
	if !orderExists {
		return models.OrderResponse{}, ErrOrderNotExists
	}
	return models.OrderResponse{Order: order, Items: orderItems}, nil

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

	query := "INSERT INTO orders(" + strings.Join(columns, ", ") + ") VALUES(" + strings.Join(placeholders, ", ") + ") RETURNING id, CURRENT_TIMESTAMP,status"
	var createdAt time.Time
	var orderId int
	var status string
	err = tr.QueryRow(query, args...).Scan(&orderId, &createdAt, &status)
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

	return models.OrderResponse{Order: models.Order{ID: orderId, UserID: order.UserID, Status: &status, CreatedAt: createdAt}, Items: OrderItems}, nil

}
