package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// CORSMiddleware returns echo.handlerFunc middleware to handle CORS
func CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if ctx.Request().Method == http.MethodOptions {
			err := ctx.NoContent(http.StatusNoContent)
			return err
		}
		return next(ctx)
	}
}