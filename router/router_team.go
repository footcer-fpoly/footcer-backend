package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func TeamRouter(e *echo.Echo, sql *db.Sql) {

	teamHandler := handler.TeamHandler{
		TeamRepo: repo.NewTeamRepo(sql),
	}

	e.POST("/team/add", teamHandler.AddTeam, middleware.JWTMiddleware())
	e.POST("/team/search-phone", teamHandler.SearchWithPhone, middleware.JWTMiddleware())
	e.POST("/team/add-member", teamHandler.AddMemberTeam, middleware.JWTMiddleware())
}
