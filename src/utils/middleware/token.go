package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/utils/security"
)

// AuthMiddleware returns gin.handlerFunc middleware to handle Auth
func AuthMiddleware() gin.HandlerFunc {
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