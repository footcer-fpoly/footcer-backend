package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	"footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func NotificationRouter(e *echo.Echo, sql *db.Sql) {

	notificationHandler := handler.NotificationHandler{
		NotificationRepo: repo_impl.NewNotificationRepo(sql),
	}

	e.GET("/notification/get", notificationHandler.GetNotification, middleware.JWTMiddleware())
}
