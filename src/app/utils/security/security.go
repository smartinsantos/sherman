package security

import (
	"github.com/labstack/echo/v4"
	"sherman/src/domain/auth"
	"sync"
)

type (
	// SecurityService security.SecurityService interface definition
	SecurityService interface {
		// password
		Hash(password string) ([]byte, error)
		VerifyPassword(hashedPassword, password string) error
		// token
		GenToken(userID, tokenType string, exp int64) (string, error)
		GetAndValidateAccessToken(ctx echo.Context) (auth.TokenMetadata, error)
		GetAndValidateRefreshToken(ctx echo.Context) (auth.TokenMetadata, error)
	}

	service struct{}
)

var (
	instance SecurityService
	once     sync.Once
)

// Get returns an instance of security.SecurityService
func Get() SecurityService {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}
