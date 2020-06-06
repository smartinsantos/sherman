package middleware

import (
	"github.com/labstack/echo/v4"
	"sync"
)

type (
	// MiddlewareService middleware.MiddlewareService interface definition
	MiddlewareService interface {
		UserAuthMiddleware() echo.MiddlewareFunc
		ZeroLog() echo.MiddlewareFunc
	}

	service struct{}
)

var (
	instance MiddlewareService
	once     sync.Once
)

// Get returns an instance of middleware.MiddlewareService
func Get() MiddlewareService {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}
