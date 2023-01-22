package security

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword create a hash for a password
func HashPassword(password string) (string, error) {
	if len(password) <= 8 {
		return "", errors.New("invalid password")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(h), err
}

// PasswordValid check if hash and password matches
func PasswordValid(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
