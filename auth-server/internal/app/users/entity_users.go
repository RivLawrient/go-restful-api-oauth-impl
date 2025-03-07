package users

import "time"

type Users struct {
	Id        string
	Username  string
	Email     string
	Password  string
	BirthDate *time.Time
	CreatedAt time.Time
}
