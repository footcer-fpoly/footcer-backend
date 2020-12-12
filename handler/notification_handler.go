package handler

import (
	"footcer-backend/model"
	"footcer-backend/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

type NotificationHandler struct {
	NotificationRepo repository.NotificationRepository
}

func (n *NotificationHandler) AddNotification(c echo.Context) error {
	return nil
}

func (n *NotificationHandler) GetNotification(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	notify, err := n.NotificationRepo.GetNotification(c.Request().Context(), claims.UserId)
	if err != nil {

		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       notify,
	})
}

