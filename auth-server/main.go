package main

import (
	"auth-server/internal/config"
	"auth-server/internal/router"

	"github.com/labstack/echo/v4"
)

func main() {
	ec := echo.New()
	router.Routes(ec)
	ec.GET("/", func(c echo.Context) error {
		return c.String(200, "auth server")
	})

	server := config.NewServer(ec)

	server.ListenAndServe()
}
