package handler

import (
	"footcer-backend/repository"
	"github.com/labstack/echo"
)

type NotificationHandler struct {
	NotificationRepo repository.NotificationRepository
}

func (n *NotificationHandler) AddNotification(c echo.Context) error {
return  nil
}


func (n *NotificationHandler) GetNotification(c echo.Context) error {
	return  nil
}