package handler

import (
	"footcer-backend/helper"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/security"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (u *UserHandler) Profile(c echo.Context) error {
	return nil

}

func (u *UserHandler) List(c echo.Context) error {
	return nil

}
func (u *UserHandler) Create(c echo.Context) error {
	req := model.User{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.UserId = uuid.NewV1().String()
	req.Role = 0

	user, err := u.UserRepo.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	token, err := security.GenToken(user)
	if err != nil {
		//log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}
func (u *UserHandler) Update(c echo.Context) error {
	return nil

}
func (u *UserHandler) CheckValidPhone(c echo.Context) error {
	req := model.User{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	valid := u.UserRepo.ValidPhone(c.Request().Context(), req.Phone)
	if valid != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    "Phone is valid",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusConflict, model.Response{
		StatusCode: http.StatusConflict,
		Message:    "Phone exits",
		Data:       nil,
	})

}
func (u *UserHandler) CreateForPhone(c echo.Context) error {
	req := model.User{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.UserId = uuid.NewV1().String()

	req.Role = 0
	hash := security.HashAndSalt([]byte(req.Password))
	req.Password = hash

	user, err := u.UserRepo.CreateForPhone(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req := model.ReqSignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}

	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}
