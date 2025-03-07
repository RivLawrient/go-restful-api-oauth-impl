package router

import (
	"auth-server/internal/app/users"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo) {
	app.POST("/users", users.RegisterUsersHandle)
	app.PUT("/users", users.LoginUsersHandle)
}
