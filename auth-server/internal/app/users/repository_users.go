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

func FindByEmail(email string) (*Users, error) {
	log.Println("FindByEmailUsersRepo")

	query := "SELECT id, username,password, birth_date, created_at FROM users WHERE email=$1"
	result := config.GetConnection().QueryRow(query, &email)

	if result.Err() != nil {
		return nil, result.Err()
	}

	data := new(Users)
	if err := result.Scan(&data.Id, &data.Username, &data.Password, &data.BirthDate, &data.CreatedAt); err != nil {
		return nil, err
	}

	data.Email = email

	return data, nil
}

func DeleteAll() error {
	log.Println("DeleteAllUsersRepo")

	query := "TRUNCATE users"
	_, err := config.GetConnection().Exec(query)

	return err
}
