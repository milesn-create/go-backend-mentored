package models

import "time"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
}
type UserUpdate struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Status    *string   `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type OrderItem struct {
	ID          int    `json:"id"`
	OrderId     int    `json:"order_id"`
	ProductName string `json:"product_name"`
	Price       *int   `json:"price"`
	Quantity    int    `json:"quantity"`
}
type CreateOrderItemRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	Price       *int   `json:"price" binding:"required,gt=-1"`
	Quantity    int    `json:"quantity" binding:"required,gt=0"`
}
type CreateOrderRequest struct {
	UserID     int                      `json:"user_id" binding:"required"`
	Status     *string                  `json:"status"`
	OrderItems []CreateOrderItemRequest `json:"order_items" binding:"required,min=1,dive"`
}
type OrderResponse struct {
	Order
	Items []OrderItem `json:"items"`
}
