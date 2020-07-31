package handler

import (
	"footcer-backend/helper"
	"footcer-backend/model"
	"footcer-backend/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type OrderHandler struct {
	OrderRepo repository.OrderRepository
}

func (o *OrderHandler) AddOrder(c echo.Context) error {
	req := model.Order{}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.OrderId = uuid.NewV1().String()
	req.UserId = claims.UserId
	req.Accept = "0"
	req.Finish = "0"
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	teamDetails, err := o.OrderRepo.AddOrder(c.Request().Context(), req)
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
		Data:       teamDetails,
	})

}

func (o *OrderHandler) AcceptOrder(c echo.Context) error {
	req := model.Order{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	req.Accept = "1"

	err := o.OrderRepo.AcceptOrder(c.Request().Context(), req)
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
		Data:       nil,
	})

}

func (o *OrderHandler) RefuseOrder(c echo.Context) error {
	req := model.Order{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	req.Accept = "-1"

	err := o.OrderRepo.AcceptOrder(c.Request().Context(), req)
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
		Data:       nil,
	})

}

func (o *OrderHandler) FinishOrder(c echo.Context) error {
	req := model.Order{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	req.Finish = "1"

	err := o.OrderRepo.FinishOrder(c.Request().Context(), req)
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
		Data:       nil,
	})

}

func (o *OrderHandler) ListOrderForStadium(c echo.Context) error {
	stadiumID := c.Param("id")

	orders, err := o.OrderRepo.ListOrderForStadium(c.Request().Context(), stadiumID)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       err.Error,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       orders,
	})
}

func (o *OrderHandler) ListOrderForUser(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	orders, err := o.OrderRepo.ListOrderForUser(c.Request().Context(), claims.UserId)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       err.Error,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       orders,
	})
}
