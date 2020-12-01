package handler

import (
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"github.com/labstack/echo"
	"net/http"
)

type AdminHandler struct {
	AdminRepo repository.AdminRepository
}

func (a *AdminHandler) AcceptStadium(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	id := c.Param("id")

	err := a.AdminRepo.AcceptStadium(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    message.Success,
		Data:       nil,
	})
}

func (a *AdminHandler) Statistics(c echo.Context) error {

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")


	statistics, err := a.AdminRepo.Statistics(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    message.Success,
		Data:       statistics,
	})
}
