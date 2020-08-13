package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func AdRouter(e *echo.Echo, sql *db.Sql) {

	adHandler := handler.AdHandler{
		AdRepo: repo.NewAdRepo(sql),
	}

	e.POST("/ad/add", adHandler.AddAd, middleware.JWTMiddleware())
}
