package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sherman/src/app/utils/response"
)

// UserAuthMiddleware returns echo.HandlerFunc middleware to handle user auth
func (s *service) JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if _, err := s.securityService.GetAndValidateAccessToken(ctx); err != nil {
				res := response.NewResponse()
				res.SetError(http.StatusUnauthorized, "invalid token")
				return ctx.JSON(http.StatusUnauthorized, res.GetBody())
			}
			return next(ctx)
		}
	}
}
