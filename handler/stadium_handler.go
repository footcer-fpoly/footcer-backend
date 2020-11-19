package handler

import (
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/upload"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type StadiumHandler struct {
	StadiumRepo repository.StadiumRepository
}

func (u *StadiumHandler) StadiumInfo(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	stadium, err := u.StadiumRepo.StadiumInfo(c.Request().Context(), claims.UserId)
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
		Data:       stadium,
	})
}

func (u *StadiumHandler) StadiumInfoForID(c echo.Context) error {
	stadiumID := c.Param("id")

	stadium, err := u.StadiumRepo.StadiumInfoForID(c.Request().Context(), stadiumID)
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
		Data:       stadium,
	})
}

func (u *StadiumHandler) UpdateStadium(c echo.Context) error {
	req := model.Stadium{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	urls, errUpload := upload.Upload(c)
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}
	image := ""

	if len(urls) > 0 {
		image = urls[0]
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	stadium := model.Stadium{
		StadiumName: req.StadiumName,
		Address:     req.Address,
		Description: req.Description,
		Image:       image,

		Category:  req.Category,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Ward:      req.Ward,
		District:  req.District,
		City:      req.City,
		UserId:    claims.UserId,
	}

	stadium, err := u.StadiumRepo.StadiumUpdate(c.Request().Context(), stadium, claims.Role)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (u *StadiumHandler) UpdateStadiumCollage(c echo.Context) error {
	req := model.StadiumCollage{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	stadiumColl := model.StadiumCollage{
		StadiumCollageId:   req.StadiumCollageId,
		NameStadiumCollage: req.NameStadiumCollage,
		AmountPeople:       req.AmountPeople,
		StartTime:          req.StartTime,
		EndTime:            req.EndTime,
		PlayTime:           req.PlayTime,
		StadiumId:          req.StadiumId,
		DefaultPrice:       req.DefaultPrice,
	}

	stadiumColl, err := u.StadiumRepo.StadiumCollageUpdate(c.Request().Context(), stadiumColl)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (u *StadiumHandler) AddStadiumCollage(c echo.Context) error {
	req := model.StadiumCollage{}
	req.StadiumCollageId = uuid.NewV1().String()

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	user, err := u.StadiumRepo.StadiumCollageAdd(c.Request().Context(), req)
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
		Data:       user,
	})
}

func (u *StadiumHandler) SearchStadiumLocation(c echo.Context) error {
	latitude := c.QueryParam("latitude")
	longitude := c.QueryParam("longitude")
	stadium, err := u.StadiumRepo.SearchStadiumLocation(c.Request().Context(), latitude, longitude)
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
		Data:       stadium,
	})

}

func (u *StadiumHandler) SearchStadiumName(c echo.Context) error {
	name := c.Param("name")

	stadium, err := u.StadiumRepo.SearchStadiumName(c.Request().Context(), name)
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
		Data:       stadium,
	})
}

func (u *StadiumHandler) StadiumDetailsInfoForStadiumCollage(c echo.Context) error {
	id := c.Param("id")

	stadiumDet, err := u.StadiumRepo.StadiumDetailsInfoForStadiumCollage(c.Request().Context(), id)
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
		Data:       stadiumDet,
	})
}

func (u *StadiumHandler) StadiumDetailsUpdateForStadiumCollage(c echo.Context) error {
	req := model.StadiumDetails{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	_, err := u.StadiumRepo.StadiumDetailsUpdateForStadiumCollage(c.Request().Context(), req)
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

func (u *StadiumHandler) StadiumCollageDelete(c echo.Context) error {
	idCollage := c.Param("id")

	err := u.StadiumRepo.StadiumCollageDelete(c.Request().Context(), idCollage)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lí thành công",
		Data:       nil,
	})
}

func (u *StadiumHandler) StadiumUploadImages(c echo.Context) error {
	req := model.Images{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	urls, errUpload := upload.Upload(c)
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}

	if len(urls) > 0 {
		for i := 0; i < len(urls); i++ {
			req.ImageId = uuid.NewV4().String()
			req.Url = urls[i]
			_, err := u.StadiumRepo.StadiumUploadImages(c.Request().Context(), req)
			if err != nil {
				return c.JSON(http.StatusOK, model.Response{
					StatusCode: http.StatusConflict,
					Message:    err.Error(),
				})
			}
		}
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Thất bại",
		Data:       nil,
	})
}
