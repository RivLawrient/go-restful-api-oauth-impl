package main

import "github.com/labstack/echo/v4"

func main() {
	ec := echo.New()
	ec.GET("/", func(c echo.Context) error {
		return c.String(200, "show server")
	})

	ec.Logger.Fatal(ec.Start(":3001"))
}
