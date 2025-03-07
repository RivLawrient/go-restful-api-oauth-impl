package helper

import (
	"log"
	"regexp"
)

func EmailValidation(email string) bool {
	log.Println("validating email format")
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	log.Println(re.MatchString(email))
	return re.MatchString(email)
}
