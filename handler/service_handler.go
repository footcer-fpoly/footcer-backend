package handler

import (
	"footcer-backend/helper"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/upload"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type ServiceHandler struct {
	ServiceRepo repository.ServiceRepository
}

func (s *ServiceHandler) AddService(c echo.Context) error {
	req := model.Service{}
	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	urls, errUpload := upload.Upload(c)
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}

	req.ServiceId = uuid.NewV1().String()
	req.Image = urls[0]

	service, err := s.ServiceRepo.AddService(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       service,
	})
}

func (s *ServiceHandler) DeleteService(c echo.Context) error {
	err := s.ServiceRepo.DeleteService(c.Request().Context(), c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}
