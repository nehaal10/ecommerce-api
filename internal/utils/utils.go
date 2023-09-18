package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PasswordHash(password string) string {
	datapass, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	Checkerr(err)

	return string(datapass)
}

func ComaparePass(hashed string, pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass)); err != nil {
		return false
	}

	return true

}
