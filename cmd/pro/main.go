package main

import (
	"footcer-backend/db"
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/model"
	"footcer-backend/router"
	"footcer-backend/security/pro"
	"github.com/labstack/echo"
)

func init() {
	//os.Setenv("APP_NAME", "footcer")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     pro.HOST,
		Port:     pro.PORT,
		UserName: pro.USERNAME,
		Password: pro.PASSWORD,
		DbName:   pro.DB_NAME,
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
			Message:    "This is the website Footer Team :=))",
		})
	})
	router.UserRouter(e, sql)
	router.StadiumRouter(e, sql)
	router.ReviewRouter(e, sql)
	router.ServiceRouter(e, sql)
	router.OrderRouter(e, sql)
	router.TeamRouter(e, sql)
	router.GameRouter(e, sql)
	router.NotificationRouter(e, sql)
	router.AdRouter(e, sql)

	//upload
	e.Static("/static", "../images/")

	e.Logger.Fatal(e.Start(":4000"))

}
