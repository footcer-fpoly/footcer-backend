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

type ReviewHandler struct {
	ReviewRepo repository.ReviewRepository
}

func (u *ReviewHandler) AddReview(c echo.Context) error {
	req := model.Review{}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	req.ReviewId = uuid.NewV1().String()
	req.UserId = claims.UserId
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	req.User.UserId = claims.UserId

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	_, err := u.ReviewRepo.AddReview(c.Request().Context(), req)
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
