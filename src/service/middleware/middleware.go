package middleware

import (
	"github.com/labstack/echo/v4"
	"sherman/src/app/config"
	cmc "sherman/src/service/middleware/config"
	"sherman/src/service/security"
)

type (
	// Middleware middleware.Middleware interface definition
	Middleware interface {
		JWT() echo.MiddlewareFunc
		ZeroLog() echo.MiddlewareFunc
		ZeroLogWithConfig(cfg *cmc.ZeroLogConfig) echo.MiddlewareFunc
	}

	service struct {
		config          *config.GlobalConfig
		securityService security.Security
	}
)

// New returns an instance of middleware.Middleware
func New(cfg *config.GlobalConfig, ss security.Security) Middleware {
	return &service{
		config:          cfg,
		securityService: ss,
	}
}
