package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) string {

	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatalf("Don't hashed the password: %v", err)
	}

	return string(hashedPassword)
}
