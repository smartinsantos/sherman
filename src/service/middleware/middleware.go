package middleware

import (
	"github.com/labstack/echo/v4"
	"sherman/src/service/security"
)

type (
	// Middleware middleware.Middleware interface definition
	Middleware interface {
		JWT() echo.MiddlewareFunc
		ZeroLog() echo.MiddlewareFunc
	}

	service struct {
		securityService security.Security
	}
)

// New returns an instance of middleware.Middleware
func New(ss security.Security) Middleware {
	return &service{
		securityService: ss,
	}
}
