package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	"github.com/labstack/echo"
	repo "footcer-backend/repository/repo_impl"

)

func GameRouter(e *echo.Echo, sql *db.Sql) {

	gameHandler := handler.GameHandler{
		GameRepo: repo.NewGameRepo(sql),
	}
	e.POST("/game/add", gameHandler.AddGame, middleware.JWTMiddleware())

}