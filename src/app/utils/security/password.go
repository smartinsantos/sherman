package security

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	PasswordUtil interface {
		Hash(password string) ([]byte, error)
		VerifyPassword(hashedPassword, password string) error
	}

	passwordUtil struct{}
)

// NewPasswordUtil PasswordUtil constructor
func NewPasswordUtil() PasswordUtil {
	return &passwordUtil{}
}

// Hash hashes a script with bcrypt
func (pu *passwordUtil) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword un-hashes a string from a bcrypt hash
func (pu *passwordUtil) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
