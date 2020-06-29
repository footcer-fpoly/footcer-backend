package handler

import (
	"footcer-backend/model"
	"footcer-backend/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

type StadiumHandler struct {
	StadiumRepo repository.StadiumRepository
}

func (u *StadiumHandler) StadiumInfo(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	stadium, err := u.StadiumRepo.StadiumInfo(c.Request().Context(), claims.UserId)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       stadium,
	})
}
