package main

import (
	"footcer-backend/db"
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/model"
	"footcer-backend/router"
	"footcer-backend/security"
	"github.com/labstack/echo"
)

func init() {
	//os.Setenv("APP_NAME", "footcer")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     security.HOST,
		Port:     security.PORT,
		UserName: security.USERNAME,
		Password: security.PASSWORD,
		DbName:   security.DB_NAME,
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, model.Response{
			StatusCode: 200,
			Message:    "Home Page",
		})
	})
	router.UserRouter(e, sql)
	router.StadiumRouter(e, sql)
	router.ReviewRouter(e, sql)
	router.ServiceRouter(e, sql)
	router.OrderRouter(e, sql)
	router.TeamRouter(e, sql)
	router.GameRouter(e, sql)

	//upload
	e.Static("/static", "../images/")

	e.Logger.Fatal(e.Start(":4000"))

}
