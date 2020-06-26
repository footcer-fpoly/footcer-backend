package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func UserRouter(e *echo.Echo, sql *db.Sql) {

	userHandler := handler.UserHandler{
		UserRepo: repo.NewUserRepo(sql),
	}

	e.POST("/users/sign-in", userHandler.Create)
	e.POST("/users/valid-phone", userHandler.CheckValidPhone)
	e.POST("/users/sign-up-phone", userHandler.CreateForPhone)
	e.POST("/users/sign-in-phone", userHandler.HandleSignIn)
	//e.PUT("/user/update", userHandler.Update,middleware.JWTMiddleware())

	//e.GET("/user/profile", userHandler.Profile, middleware.JWTMiddleware())
	//e.GET("/user/list", userHandler.List, middleware.JWTMiddleware())
}
