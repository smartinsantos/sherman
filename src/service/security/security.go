package security

import (
	"github.com/labstack/echo/v4"
	"sherman/src/app/config"
	"sherman/src/domain/auth"
)

type (
	// Security security.Security interface definition
	Security interface {
		// password
		Hash(password string) ([]byte, error)
		VerifyPassword(hashedPassword, password string) error
		// token
		GenToken(userID, tokenType string, iat, exp int64) (string, error)
		GetAndValidateAccessToken(ctx echo.Context) (auth.TokenMetadata, error)
		GetAndValidateRefreshToken(ctx echo.Context) (auth.TokenMetadata, error)
	}

	service struct {
		config *config.GlobalConfig
	}
)

// New returns an instance of security.Security
func New(cfg *config.GlobalConfig) Security {
	return &service{
		config: cfg,
	}
}
