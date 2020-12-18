package main

import (
	"footcer-backend/db"
	"footcer-backend/helper"
	"footcer-backend/log"
	//"footcer-backend/model"
	"footcer-backend/router"
	"footcer-backend/security/pro"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

func init() {
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
func PrivacyPolicy(c echo.Context) error {
	return c.Render(http.StatusOK, "privacy-policy.html", "Footcer")
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

	t := &Template{
		templates: template.Must(template.ParseGlob("../../public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", Web)
	e.GET("/privacy-policy", PrivacyPolicy)

	router.UserRouter(e, sql)
	router.StadiumRouter(e, sql)
	router.ReviewRouter(e, sql)
	router.ServiceRouter(e, sql)
	router.OrderRouter(e, sql)
	router.TeamRouter(e, sql)
	router.GameRouter(e, sql)
	router.NotificationRouter(e, sql)
	router.StatisticsRouter(e, sql)
	router.AdminRouter(e, sql)

	//upload
	e.Static("/static", "../../images/")

	e.Logger.Fatal(e.Start(":4000"))

}
