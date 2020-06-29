package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func StadiumRouter(e *echo.Echo, sql *db.Sql) {
	stadiumHandler := handler.StadiumHandler{
		StadiumRepo: repo.NewStadiumRepo(sql),
	}
	e.GET("/stadium/info", stadiumHandler.StadiumInfo, middleware.JWTMiddleware())
}
