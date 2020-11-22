package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func StatisticsRouter(e *echo.Echo, sql *db.Sql) {

	statisticsHandler := handler.StatisticsHandler{
		StatisticsRepo: repo.NewStatisticsRepo(sql),
	}

	e.GET("/statistics", statisticsHandler.Statistics, middleware.JWTMiddleware())
}
