package handler

import (
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/model"
	"footcer-backend/repository"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type GameHandler struct {
	GameRepo repository.GameRepository
}
func (g *GameHandler) AddGame (c echo.Context) error{
	req := model.Game{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.GameId = uuid.NewV1().String()
	req.Finish = "0"
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	game, err := g.GameRepo.AddGame(c.Request().Context(), req)
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
		Data:       game,
	})

}


func (g * GameHandler) JoinGame(c echo.Context) error{
	req := model.GameTemp{}

	req.GameTempId = uuid.NewV1().String()
	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.JoinGame(c.Request().Context(), req)
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


func (g * GameHandler) AcceptJoin(c echo.Context) error{
	req := model.GameTemp{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.AcceptJoin(c.Request().Context(), req)
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

func (g * GameHandler) RefuseJoin(c echo.Context) error{
	req := model.GameTemp{}

	defer c.Request().Body.Close()
	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	err := g.GameRepo.RefuseJoin(c.Request().Context(), req)
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

func (g * GameHandler) GetGame(c echo.Context) error{
	date := c.Param("date")

	games,err := g.GameRepo.GetGame(c.Request().Context(), date)
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
		Data:       games,
	})
}
