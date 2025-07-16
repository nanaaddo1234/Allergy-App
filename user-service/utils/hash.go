package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plaintext password and returns a bcrypt-hashed version.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a plaintext password with a hashed password and returns true if they match.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
