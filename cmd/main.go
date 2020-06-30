package main

import (
	"footcer-backend/db"
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/model"
	"footcer-backend/router"
	"github.com/labstack/echo"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "footcer")
	log.InitLogger(true)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "123456",
		DbName:   "footcerdb",
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

	e.Logger.Fatal(e.Start(":4000"))

}
