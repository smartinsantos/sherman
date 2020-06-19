package config

import (
	emw "github.com/labstack/echo/v4/middleware"
	"net/http"
)

// DefaultCorsConfig is the default ZeroLog middleware config.
var CustomCorsConfig = emw.CORSConfig{
	AllowOrigins:     []string{"*"},
	AllowCredentials: true,
	AllowHeaders: []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"accept",
		"origin",
		"Cache-Control",
		"X-Requested-With"},
	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
}
