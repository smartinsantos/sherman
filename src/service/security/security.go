package security

import (
	"github.com/labstack/echo/v4"
	"sherman/src/domain/auth"
)

type (
	// Security security.Security interface definition
	Security interface {
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

// New returns an instance of security.Security
func New() Security {
	return &service{}
}
