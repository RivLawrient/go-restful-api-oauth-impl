package config

import (
	"auth-server/internal/app/users"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo) {
	app.GET("/user", users.GetUsersHanle)
}
