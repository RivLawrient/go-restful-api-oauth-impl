package helper

import (
	"log"
	"regexp"
)

func EmailValidation(email string) bool {
	log.Println("validating email format")
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return re.MatchString(email)
}
