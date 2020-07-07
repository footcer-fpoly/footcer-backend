package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func ReviewRouter(e *echo.Echo, sql *db.Sql) {

	reviewHandler := handler.ReviewHandler{
		ReviewRepo: repo.NewReviewRepo(sql),
	}

	e.POST("/review/add", reviewHandler.AddReview, middleware.JWTMiddleware())
}
