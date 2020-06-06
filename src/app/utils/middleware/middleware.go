package middleware

import (
	"github.com/labstack/echo/v4"
	"sync"
)

type (
	// Service middleware.Service interface definition
	Service interface {
		UserAuthMiddleware() echo.MiddlewareFunc
		ZeroLog() echo.MiddlewareFunc
	}

	service struct{}
)

var (
	instance Service
	once     sync.Once
)

// Get returns an instance of middleware.Service
func Get() Service {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}
