package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func AdminRouter(e *echo.Echo, sql *db.Sql) {

	adminHandler := handler.AdminHandler{
		AdminRepo: repo.NewAdminRepo(sql),
	}

	e.GET("/admin/accept-stadium/:id", adminHandler.AcceptStadium)
	e.GET("/admin/statistics", adminHandler.Statistics)

}
