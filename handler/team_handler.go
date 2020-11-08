package handler

import (
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/upload"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type TeamHandler struct {
	TeamRepo repository.TeamRepository
}

func (t *TeamHandler) AddTeam(c echo.Context) error {
	req := model.Team{}

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	avatar := ""
	background := ""

	urls, errUpload := upload.UploadForKey(c, "avatar")
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}
	if len(urls) > 0 {
		if urls[0] != "" {
			avatar = urls[0]
		} else {
			avatar = "http://footcer.tk:4000/static/team/example_avatar_team.jpg"
		}
	}

	urls, errUpload = upload.UploadForKey(c, "background")
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}
	if len(urls) > 0 {
		if urls[0] != "" {
			background = urls[0]
		} else {
			background = "http://footcer.tk:4000/static/team/example_background_team.jpg"
		}
	}

	req = model.Team{
		TeamId:      uuid.NewV1().String(),
		Name:        req.Name,
		Place:       req.Place,
		Description: req.Description,
		Avatar:      avatar,
		Background:  background,
		Level:       req.Level,
		LeaderId:    claims.UserId,
		MemberList:  req.MemberList,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	user, err := t.TeamRepo.AddTeam(c.Request().Context(), req)
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
		Data:       user,
	})
}

func (t *TeamHandler) SearchWithPhone(c echo.Context) error {
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

	user, err := t.TeamRepo.SearchWithPhoneMemberTeam(c.Request().Context(), req.Phone)
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
		Data:       user,
	})
}

func (t *TeamHandler) AddMemberTeam(c echo.Context) error {
	req := model.TeamDetails{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	req.TeamDetailsId = uuid.NewV1().String()
	req.Accept = "0"
	req.Role = "0"
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	teamDetails, err := t.TeamRepo.AddMemberTeam(c.Request().Context(), req, claims.UserId)
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
		Data:       teamDetails,
	})

}

func (t *TeamHandler) GetTeamForUserAccept(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	user, err := t.TeamRepo.GetTeamForUserAccept(c.Request().Context(), claims.UserId)
	if err != nil {

		return c.JSON(http.StatusOK, model.Response{
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

func (t *TeamHandler) GetTeamForUserReject(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	user, err := t.TeamRepo.GetTeamForUserReject(c.Request().Context(), claims.UserId)
	if err != nil {

		return c.JSON(http.StatusOK, model.Response{
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

func (t *TeamHandler) DeleteMember(c echo.Context) error {

	err := t.TeamRepo.DeleteMember(c.Request().Context(), c.Param("id"))
	if err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})

}

func (t *TeamHandler) DeleteTeam(c echo.Context) error {

	err := t.TeamRepo.DeleteTeam(c.Request().Context(), c.Param("id"))
	if err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})

}

func (t *TeamHandler) UpdateTeam(c echo.Context) error {
	req := model.Team{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	avatar := ""
	background := ""
	urls, errUpload := upload.UploadForKey(c, "avatar")
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}
	if len(urls) > 0 {
		if urls[0] != "" {
			avatar = urls[0]
		}
	}

	urls, errUpload = upload.UploadForKey(c, "background")
	if errUpload != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errUpload.Error(),
		})
	}
	if len(urls) > 0 {
		if urls[0] != "" {
			background = urls[0]
		}
	}

	team := model.Team{
		TeamId:      req.TeamId,
		Name:        req.Name,
		Level:       req.Level,
		Place:       req.Place,
		Description: req.Description,
		Avatar:      avatar,
		Background:  background,
	}
	team, err := t.TeamRepo.UpdateTeam(c.Request().Context(), team)
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

func (t *TeamHandler) AcceptInvite(c echo.Context) error {

	req := model.TeamDetails{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := t.TeamRepo.AcceptInvite(c.Request().Context(), req)
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
