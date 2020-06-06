package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sherman/src/app/utils/response"
	"sherman/src/app/utils/security"
)

// UserAuthMiddleware returns echo.HandlerFunc middleware to handle user auth
func UserAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if _, err := security.Get().GetAndValidateAccessToken(ctx); err != nil {
				res := response.NewResponse()
				res.SetError(http.StatusUnauthorized, "invalid token")
				return ctx.JSON(http.StatusUnauthorized, res.GetBody())
			}
			return next(ctx)
		}
	}
}
