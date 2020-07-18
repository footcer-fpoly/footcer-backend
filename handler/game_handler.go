package handler

import (
	"footcer-backend/repository"
	"github.com/labstack/echo"
)

type GameHandler struct {
	GameRepo repository.GameRepository
}
func (t *GameHandler) AddGame (c echo.Context) error{
	return nil
}