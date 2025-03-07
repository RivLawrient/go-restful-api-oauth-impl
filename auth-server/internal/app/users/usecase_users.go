package users

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUsers(req *NewUsersRequest) error {
	log.Println("RegisterUsersUsecase")

	id := uuid.New().String()
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}
	birthDate := time.UnixMilli(req.BirthDate)

	user := Users{
		Id:        id,
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(pass),
		BirthDate: &birthDate,
	}

	repo := Create(id, &user)
	return repo
}
