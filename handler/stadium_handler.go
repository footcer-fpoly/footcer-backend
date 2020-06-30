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

	stadium, err := u.StadiumRepo.StadiumInfo(c.Request().Context(), claims.UserId) // sao may doan nay khong dua vao context, chi can dua user_id vao thoi
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

func (s *StadiumHandler) UpdateStadium(c echo.Context) error {
	req := model.Stadium{}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate thông tin gửi lên
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	stadium := model.Stadium{
		StadiumName: req.StadiumName,
		Address:     req.Address,
		Description: req.Description,
		Image:       req.Image,
		PriceNormal: req.PriceNormal,
		PricePeak:   req.PricePeak,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Category:    req.Category,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
		UserId:      claims.UserId,
	}

	stadium, err = s.StadiumRepo.StadiumUpdate(c.Request().Context(), stadium)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}
