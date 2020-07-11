package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func OrderRouter(e *echo.Echo, sql *db.Sql) {

	orderHandler := handler.OrderHandler{
		OrderRepo: repo.NewOrderRepo(sql),
	}

	e.POST("/order/add", orderHandler.AddOrder, middleware.JWTMiddleware())
}
