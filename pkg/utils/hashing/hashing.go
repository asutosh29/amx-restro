package hashing

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println("Error generating hash from password")
		return ""
	}
	return string(hashedBytes)
}

func CheckPasswordFromHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
