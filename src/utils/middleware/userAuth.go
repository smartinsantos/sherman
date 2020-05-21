package middleware

import (
	"github.com/labstack/echo"
	"net/http"
	"root/src/utils/response"
	"root/src/utils/security"
)

// UserAuthMiddleware returns echo.HandlerFunc middleware to handle user auth
func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if _, err := security.GetAndValidateAccessToken(ctx); err != nil {
			res := response.NewResponse()
			res.SetError(http.StatusUnauthorized, "invalid token")
			return ctx.JSON(http.StatusUnauthorized, res.GetBody())
		}
		return next(ctx)
	}
}