package security

import "golang.org/x/crypto/bcrypt"

// HashPassword create a hash for a password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// PasswordValid check if hash and password matches
func PasswordValid(hash []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
