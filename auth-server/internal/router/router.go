package router

import (
	"auth-server/internal/app/users"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo) {
	app.GET("/hello", func(c echo.Context) error {
		return c.String(200, "hello world")
	})
	app.POST("/users", users.RegisterUsersHandle)
	app.PUT("/users", users.LoginUsersHandle)
}
