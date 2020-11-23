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
	e.PUT("/order/update-status", orderHandler.UpdateStatusOrder, middleware.JWTMiddleware())
	e.PUT("/order/finish", orderHandler.FinishOrder, middleware.JWTMiddleware())
	e.GET("/order/stadium/:id", orderHandler.ListOrderForStadium, middleware.JWTMiddleware())
	e.GET("/order/user", orderHandler.ListOrderForUser, middleware.JWTMiddleware())
	e.GET("/order/:id", orderHandler.ListOrderForUser, middleware.JWTMiddleware())

}
