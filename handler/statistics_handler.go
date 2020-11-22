package handler

import (
	"footcer-backend/model"
	"footcer-backend/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

type StatisticsHandler struct {
	StatisticsRepo repository.StatisticsRepository
}

func (s *StatisticsHandler) Statistics(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role == 0 {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    "Đây là chức năng xem thống kê chỉ dành cho chủ sân",
			Data:       nil,
		})
	}

	statistics := model.Statistics{}
	var err error
	day := c.QueryParam("day")
	month := c.QueryParam("month")
	date := ""
	message := ""
	if len(day) != 0 {
		date = day
		message = "Thống kê theo ngày " + day
		statistics, err = s.StatisticsRepo.StatisticsDay(c.Request().Context(), date, claims.UserId)
		if err != nil {
			return c.JSON(http.StatusOK, model.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}

	}
	if len(month) != 0 {
		date = month
		message = "Thống kê theo tháng " + month
		statistics, err = s.StatisticsRepo.StatisticsMonth(c.Request().Context(), date, claims.UserId)
		if err != nil {
			return c.JSON(http.StatusOK, model.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       statistics,
	})
}
