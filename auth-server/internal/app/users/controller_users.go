package users

import "github.com/labstack/echo/v4"

func GetUsersHanle(c echo.Context) error {
	return c.String(201, "users")
}
