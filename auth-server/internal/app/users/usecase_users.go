package users

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUsers(req *NewUsersRequest) (*NewUsersResponse, error) {
	log.Println("RegisterUsersUsecase")

	id := uuid.New().String()
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
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
	if repo != nil {
		return nil, repo
	}

	return &NewUsersResponse{
		Id:        id,
		Username:  req.Username,
		Email:     req.Email,
		BirthDate: birthDate.Unix(),
	}, nil
}

func LoginUsers(req *LoginUsersRequest) (*LoginUsersResponse, error) {
	log.Println("LoginUsersUsecase")

	user, err := FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if result != nil {
		return nil, errors.New("password is invalid")
	}

	key := "soem"
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"email":      user.Email,
		"birth_date": user.BirthDate,
	})
	s, _ := t.SignedString([]byte(key))

	return &LoginUsersResponse{
		TokenJwt: s,
	}, nil
}
