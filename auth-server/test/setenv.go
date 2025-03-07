package test

import "os"

func MockEnv() {
	os.Setenv("DB_USERNAME", "lawrient")
	os.Setenv("DB_NAME_AUTH", "user_auth")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
}
