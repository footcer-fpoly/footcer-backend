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
	e.POST("/game/join", gameHandler.JoinGame, middleware.JWTMiddleware())
	e.POST("/game/accept", gameHandler.AcceptJoin, middleware.JWTMiddleware())
	e.POST("/game/refuse", gameHandler.RefuseJoin, middleware.JWTMiddleware())

}