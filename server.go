package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/blog", func(c echo.Context) error {
		return c.Render(http.StatusOK, "blog.html", nil)
	})

	e.GET("/blog/:postId", func(c echo.Context) error {
		postId := c.Param("postId")
		return c.Render(http.StatusOK, "blog.html", postId)
	})

	e.Static("/", ".")

	e.Logger.Fatal(e.Start(":1323"))
}
