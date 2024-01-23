package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/crims1n/moogle/pkg/podbean"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = GetTemplates()

	info, err := podbean.GetPodbean()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", info)
	})

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))

}
