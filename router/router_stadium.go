package router

import (
	"footcer-backend/db"
	"footcer-backend/handler"
	"footcer-backend/middleware"
	repo "footcer-backend/repository/repo_impl"
	"github.com/labstack/echo"
)

func StadiumRouter(e *echo.Echo, sql *db.Sql) {
	stadiumHandler := handler.StadiumHandler{
		StadiumRepo: repo.NewStadiumRepo(sql),
	}
	e.GET("/stadium/info", stadiumHandler.StadiumInfo, middleware.JWTMiddleware())
	e.GET("/stadium/info-id/:id", stadiumHandler.StadiumInfoForID, middleware.JWTMiddleware())
	e.PUT("/stadium/update", stadiumHandler.UpdateStadium, middleware.JWTMiddleware())
	e.PUT("/stadium/update-collage", stadiumHandler.UpdateStadiumCollage, middleware.JWTMiddleware())
	e.POST("/stadium/add-collage", stadiumHandler.AddStadiumCollage, middleware.JWTMiddleware())
	e.DELETE("/stadium/delete_collage/:id", stadiumHandler.StadiumCollageDelete, middleware.JWTMiddleware())
	e.GET("/stadium/search-location", stadiumHandler.SearchStadiumLocation, middleware.JWTMiddleware())
	e.GET("/stadium/search-name/:name", stadiumHandler.SearchStadiumName, middleware.JWTMiddleware())
	e.GET("/stadium/collage-details/:id", stadiumHandler.StadiumDetailsInfoForStadiumCollage, middleware.JWTMiddleware())
	e.PUT("/stadium/collage-details-update", stadiumHandler.StadiumDetailsUpdateForStadiumCollage, middleware.JWTMiddleware())
}
