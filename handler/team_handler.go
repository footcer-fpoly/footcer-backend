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

type TeamHandler struct {
	TeamRepo repository.TeamRepository
}

func (t *TeamHandler) AddTeam(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	req := model.Team{
		TeamId:    uuid.NewV1().String(),
		LeaderId:  claims.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	user, err := t.TeamRepo.AddTeam(c.Request().Context(), req)
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

func (t *TeamHandler) SearchWithPhone(c echo.Context) error {
	req := model.User{}
	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}
	user, err := t.TeamRepo.SearchWithPhoneMemberTeam(c.Request().Context(), req.Phone)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
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

	req.TeamDetailsId = uuid.NewV1().String()
	req.Accept = "0"
	req.Accept = "0"
	req.CreatedAt =  time.Now()
	req.UpdatedAt =  time.Now()


	teamDetails, err := t.TeamRepo.AddMemberTeam(c.Request().Context(), req)
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

func (t * TeamHandler) GetTeamForUser(c echo.Context) error{
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	user, err := t.TeamRepo.GetTeamForUser(c.Request().Context(), claims.UserId)
	if err != nil {
		
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

func (t * TeamHandler) GetTeamForID(c echo.Context) error{
	teamID := c.Param("id")
	user, err := t.TeamRepo.GetTeamForID(c.Request().Context(), teamID)
	if err != nil {
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
