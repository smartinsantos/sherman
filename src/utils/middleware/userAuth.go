package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/utils/security"
)

// UserAuthMiddleware returns gin.handlerFunc middleware to handle user auth
func UserAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if _, err := security.GetAndValidateAccessToken(context.Request); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "invalid token",
			})
			return
		}
		context.Next()
	}
}