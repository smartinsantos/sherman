package security

import (
	"github.com/labstack/echo/v4"
	"sherman/src/domain/auth"
	"sync"
)

type (
	// Service security.Service interface definition
	Service interface {
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
	instance Service
	once     sync.Once
)

// Get returns an instance of security.Service
func Get() Service {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}
