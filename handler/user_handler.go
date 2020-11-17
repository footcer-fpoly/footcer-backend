package handler

import (
	"footcer-backend/helper"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/model/req"
	"footcer-backend/repository"
	"footcer-backend/security"
	"footcer-backend/upload"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (u *UserHandler) Profile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user, err := u.UserRepo.SelectById(c.Request().Context(), claims.UserId)
	if err != nil {
		if err == message.UserNotFound {
			return c.JSON(http.StatusOK, model.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
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

func (u *UserHandler) List(c echo.Context) error {
	return nil

}

func (u *UserHandler) Create(c echo.Context) error {
	req := model.User{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	print(req.UserId)
	//req.UserId = uuid.NewV1().String()
	req.Role = 0
	req.TokenNotify = ""

	user, err := u.UserRepo.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
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
	req := model.User{}
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
			StatusCode: http.StatusOK,
			Message:    errUpload.Error(),
		})
	}
	avatar := ""

	if len(urls) > 0 {
		avatar = urls[0]
	}

	// validate thông tin gửi lên

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	user := model.User{
		UserId:      claims.UserId,
		DisplayName: req.DisplayName,
		Email:       req.Email,
		Phone:       req.Phone,
		Avatar:      avatar,
		Birthday:    req.Birthday,
		Position:    req.Position,
		Level:       req.Level,
	}

	user, err := u.UserRepo.Update(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}

func (u *UserHandler) CreateForPhone(c echo.Context) error {
	req := model.User{}
	req.Folder = "default"
	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	req.UserId = uuid.NewV1().String()
	hash := security.HashAndSalt([]byte(req.Password))
	req.Password = hash
	req.Avatar = "http://footcer.tk:4000/static/user/avatar.png"
	req.TokenNotify = ""

	user, err := u.UserRepo.CreateForPhone(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
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
	req := req.ReqSignIn{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}

	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
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

func (u *UserHandler) CheckValidEmail(c echo.Context) error {
	req := model.User{}
	req.Folder = "default"

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	user, errValid := u.UserRepo.ValidEmail(c.Request().Context(), req.Email)

	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	if errValid != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    errValid.Error(),
			Data:       user.Token,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cho phép đăng ký",
		Data:       nil,
	})
}

func (u *UserHandler) CheckValidPhone(c echo.Context) error {
	req := model.User{}

	req.Folder = "default"

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
	code, user, valid := u.UserRepo.ValidPhone(c.Request().Context(), req.Phone)
	if valid != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: code,
			Message:    valid.Error(),
			Data:       user,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cho phép đăng ký",
		Data:       nil,
	})

}

func (u *UserHandler) CheckValidUUID(c echo.Context) error {

	req := model.User{}
	req.Folder = "default"

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	user, errValid := u.UserRepo.ValidUUID(c.Request().Context(), req.UserId)

	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	type customDataToken struct {
		UserId      string `json:"userId"`
		DisplayName string `json:"displayName"`
		Phone       string `json:"phone"`
		Avatar      string `json:"avatar"`
		Token       string `json:"token"`
	}

	if errValid != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    errValid.Error(),
			Data:       customDataToken{UserId: user.UserId, Phone: user.Phone, Token: token, Avatar: user.Avatar, DisplayName: user.DisplayName},
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cho phép đăng ký",
		Data:       nil,
	})

}

func (u *UserHandler) UpdatePassword(c echo.Context) error {
	req := model.User{}
	req.Folder = "default"

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	req.Password = hash

	err = u.UserRepo.UpdatePassword(c.Request().Context(), req)
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

func (u *UserHandler) DeleteUser(c echo.Context) error {
	req := model.User{}

	req.Folder = "default"

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	err = u.UserRepo.DeleteUser(c.Request().Context(), req.Phone)
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
