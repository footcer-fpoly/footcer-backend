package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func ServiceRouter(e *echo.Echo, sql *db.Sql) {

	serviceHandler := handler.ServiceHandler{
		ServiceRepo: repo.NewServiceRepo(sql),
	}

	e.POST("/service/add", serviceHandler.AddService, middleware.JWTMiddleware())
	e.DELETE("/service/delete/:id", serviceHandler.DeleteService, middleware.JWTMiddleware())
}
