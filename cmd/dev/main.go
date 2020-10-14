package main

import (
	"footcer-backend/db"
	"footcer-backend/helper"
	"footcer-backend/log"
	"footcer-backend/router"
	dev "footcer-backend/security/dev"
	"github.com/labstack/echo"
	//"footcer-backend/model"
	"html/template"
	"io"
	"net/http"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "footcer")
	log.InitLogger(false)
}

type Template struct {
templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func Web(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "Footcer")
}


func main() {
	sql := &db.Sql{
		Host:     dev.HOST,
		Port:     dev.PORT,
		UserName: dev.USERNAME,
		Password: dev.PASSWORD,
		DbName:   dev.DB_NAME,
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	//t := &Template{
	//	templates: template.Must(template.ParseGlob("../../public/views/*.html")),
	//}
	//e.Renderer = t
	//
	//e.GET("/", Web)
	router.UserRouter(e, sql)
	router.StadiumRouter(e, sql)
	router.ReviewRouter(e, sql)
	router.ServiceRouter(e, sql)
	router.OrderRouter(e, sql)
	router.TeamRouter(e, sql)
	router.GameRouter(e, sql)
	router.NotificationRouter(e, sql)

	//upload
	e.Static("/static", "../../images/")


	e.Logger.Fatal(e.Start(":4001"))

}
