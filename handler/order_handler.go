package handler

import (
	"footcer-backend/repository"
	"github.com/labstack/echo"
)

type OrderHandler struct {
	OrderRepo repository.OrderRepository
}

func (o *OrderHandler) AddOrder(c echo.Context) error {
	return  nil
}
