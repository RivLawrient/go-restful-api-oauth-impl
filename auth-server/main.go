package main

import (
	"auth-server/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	ec := echo.New()
	ec.GET("/", func(c echo.Context) error {
		return c.String(200, "auth server")
	})

	server := config.NewServer(ec)

	server.ListenAndServe()
}
