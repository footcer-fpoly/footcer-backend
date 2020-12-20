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
		UserRepo: repo.NewUserRepo(sql),
		NotifyRepo: repo.NewNotificationRepo(sql),
	}

	e.POST("/team/add", teamHandler.AddTeam, middleware.JWTMiddleware())
	e.POST("/team/search-phone", teamHandler.SearchWithPhone, middleware.JWTMiddleware())
	e.POST("/team/add-member", teamHandler.AddMemberTeam, middleware.JWTMiddleware())
	e.GET("/team/:id", teamHandler.GetTeam, middleware.JWTMiddleware())
	e.GET("/team/for-user-accept", teamHandler.GetTeamForUserAccept, middleware.JWTMiddleware())
	e.GET("/team/for-user-reject", teamHandler.GetTeamForUserReject, middleware.JWTMiddleware())
	e.DELETE("/team/delete-member", teamHandler.DeleteMember, middleware.JWTMiddleware())
	e.DELETE("/team/delete-team/:id", teamHandler.DeleteTeam, middleware.JWTMiddleware())
	e.PUT("/team/update", teamHandler.UpdateTeam, middleware.JWTMiddleware())
	e.PUT("/team/accept-invite", teamHandler.AcceptInvite, middleware.JWTMiddleware())
	e.POST("/team/cancel-invite", teamHandler.CancelInvite, middleware.JWTMiddleware())
	e.POST("/team/out-team", teamHandler.OutTeam, middleware.JWTMiddleware())

}
