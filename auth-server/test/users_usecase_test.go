package test

import (
	"auth-server/internal/app/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBefore(t *testing.T) {
	MockEnv()
	err := users.DeleteAll()
	assert.Nil(t, err)
}

func TestCreateSuccess(t *testing.T) {
	TestBefore(t)

	req := users.NewUsersRequest{
		Username:  "muhsand",
		Email:     "sandi@gmail.com",
		Password:  "random",
		BirthDate: 123123123123,
	}
	resp, err := users.RegisterUsers(&req)

	assert.Equal(t, req.Username, resp.Username)
	assert.Nil(t, err)

}

func TestLoginSuccess(t *testing.T) {
	TestBefore(t)

	req := users.NewUsersRequest{
		Username:  "muhsand",
		Email:     "sandis@gmail.com",
		Password:  "12345678",
		BirthDate: 123123123123,
	}
	resp, err := users.RegisterUsers(&req)
	assert.Nil(t, err)
	assert.Equal(t, req.Username, resp.Username)

	reql := users.LoginUsersRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	login, err := users.LoginUsers(&reql)
	assert.NotNil(t, login)
	assert.Nil(t, err)
}
