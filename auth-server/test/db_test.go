package test

import (
	"auth-server/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDB(t *testing.T) {
	MockEnv()
	db := config.GetConnection()
	defer db.Close()
	err := db.Ping()

	assert.Nil(t, err)
}
