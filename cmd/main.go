package main

import (
	"footcer-backend/db"
	"footcer-backend/model"
	"github.com/labstack/echo"
)

func main()  {
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

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, model.Response{
			StatusCode: 200,
			Message:    "Home Page",
		})
	})
	e.Logger.Fatal(e.Start(":4000"))

}
