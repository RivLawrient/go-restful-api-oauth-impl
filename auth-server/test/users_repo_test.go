package test

import (
	"auth-server/internal/app/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSuccess(t *testing.T) {
	MockEnv()
	req := users.NewUsersRequest{
		Username:  "muhsand",
		Email:     "sandi@gmail.com",
		Password:  "random",
		BirthDate: 123123123123,
	}
	err := users.RegisterUsers(&req)

	assert.Nil(t, err)

}
