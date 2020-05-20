package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/utils/security"
)

// UserAuthMiddleware returns gin.handlerFunc middleware to handle user auth
func UserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, err := security.GetAndValidateAccessToken(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "invalid token",
			})
			return
		}
		ctx.Next()
	}
}