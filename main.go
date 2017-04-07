package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Any("/wssever", func(c echo.Context) error { return nil })
	e.Start(":3000")
}
