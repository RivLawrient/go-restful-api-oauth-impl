package users

import (
	"auth-server/internal/helper"
	"auth-server/internal/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUsersHandle(c echo.Context) error {
	log.Println("RegisterUsersController")
	req := new(NewUsersRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: err.Error(),
		})
	}

	if req.Username == "" {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "username field cannot be empty",
		})
	}

	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "email field cannot be empty",
		})
	}

	if !helper.EmailValidation(req.Email) {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "email format is invalid",
		})
	}

	if req.Password == "" {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "password field cannot be empty",
		})
	}

	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "password length cannot less than 8 character",
		})
	}

	if req.BirthDate == 0 {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "birth_date field cannot be empty",
		})
	}

	response, err := RegisterUsers(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.WebResponse[any]{
			Errors: err.Error(),
		})
	}

	return c.JSON(200, response)
}

func LoginUsersHandle(c echo.Context) error {
	log.Println("LoginUsersHandleController")
	req := new(LoginUsersRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: err.Error(),
		})
	}

	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "email field cannot be empty",
		})
	}

	if req.Password == "" {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "password field cannot be empty",
		})
	}

	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: "password length cannot less than 8 character",
		})
	}

	response, err := LoginUsers(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.WebResponse[any]{
			Errors: err.Error(),
		})
	}

	return c.JSON(200, model.WebResponse[LoginUsersResponse]{
		Data: *response,
	})
}
