package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash hashes a script with bcrypt
func (s *service) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword un-hashes a string from a bcrypt hash
func (s *service) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
