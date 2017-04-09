package main

import (
	"bridge/wsserver"
	"html/template"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	ws := wsserver.New()

	e.Static("/static", "./static")
	e.GET("/", func(c echo.Context) error {

		tmpl, err := template.ParseFiles("./templates/index.html")
		tmpl.Execute(c.Response(), nil)

		return err
	})
	e.Any("/wsserver/:id", ws.EchoHandler)
	e.Start(":3000")
}
