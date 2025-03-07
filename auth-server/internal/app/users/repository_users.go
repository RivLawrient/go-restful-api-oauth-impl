package users

import (
	"auth-server/internal/config"
	"log"
)

func Create(id string, req *Users) error {
	log.Println("CreateUsersRepo")

	query := "INSERT INTO users(id, username, email, password, birth_date) VALUES($1, $2, $3, $4, $5)"
	_, err := config.GetConnection().Exec(query, &id, &req.Username, &req.Email, &req.Password, &req.BirthDate)

	return err
}
