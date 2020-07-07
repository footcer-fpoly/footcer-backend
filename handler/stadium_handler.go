package handler

import (
	"footcer-backend/helper"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/upload"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
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

func (s *StadiumHandler) UpdateStadium(c echo.Context) error {
	urls, errUpload := upload.Upload(c)
	if errUpload != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}
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
		Image:       urls[0],
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

func (s *StadiumHandler) UpdateStadiumCollage(c echo.Context) error {
	req := model.StadiumCollage{}

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

	stadiumColl := model.StadiumCollage{
		StadiumCollageId:   req.StadiumCollageId,
		NameStadiumCollage: req.NameStadiumCollage,
		AmountPeople:       req.AmountPeople,
	}

	stadiumColl, err = s.StadiumRepo.StadiumCollageUpdate(c.Request().Context(), stadiumColl)
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

func (s *StadiumHandler) AddStadiumCollage(c echo.Context) error {
	req := model.StadiumCollage{}
	req.StadiumCollageId = uuid.NewV1().String()

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	user, err := s.StadiumRepo.StadiumCollageAdd(c.Request().Context(), req)
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
		Data:       user,
	})
}
