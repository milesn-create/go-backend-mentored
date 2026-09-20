package handler

import (
	"errors"
	"net/http"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/repository"
	"rest-api-gin/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var OrderService *service.OrderService

func CreateOrderHandler(c *gin.Context) {
	var RequestOrder models.CreateOrderRequest
	err := c.ShouldBindJSON(&RequestOrder)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request\n")
		return
	}
	ResponseOrder, err := OrderService.CreateOrder(RequestOrder)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseOrder)
}
func GetOrderByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "id - not an integer")
		return
	}
	orderResponse, err := OrderService.GetOrderById(id)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotExists) {
			c.String(http.StatusNotFound, repository.ErrOrderNotExists.Error())
			return
		}
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, orderResponse)
}
