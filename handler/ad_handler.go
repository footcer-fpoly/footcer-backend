package handler

import (
	"footcer-backend/repository"
	"github.com/labstack/echo"
)

type AdHandler struct {
	AdRepo repository.AdRepository
}

func (a *AdHandler) AddAd(c echo.Context) error {
	return nil
}
