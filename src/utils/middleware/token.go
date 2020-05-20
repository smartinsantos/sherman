package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"root/config"
	"root/src/domain/auth"
	"strings"
)

// extractAccessTokenMetadata extracts metadata of *jwt.Token
func extractAccessTokenMetadata(token *jwt.Token) (auth.TokenMetadata, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return auth.TokenMetadata{}, errors.New("invalid token data")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return auth.TokenMetadata{}, errors.New("invalid token data")
	}

	tokenType, ok := claims["type"].(string)
	if !ok {
		return auth.TokenMetadata{}, errors.New("invalid token data")
	}

	return auth.TokenMetadata{
		UserID: userID,
		Type: tokenType,
	}, nil
}

// getAndValidateAccessTokenFromRequest gets the access token from an http request and verifies signature
func getAndValidateAccessToken(req *http.Request) (auth.TokenMetadata, error) {
	bearToken := req.Header.Get("Authorization")
	tokenArr := strings.Split(bearToken, " ")
	if len(tokenArr) != 2 {
		return auth.TokenMetadata{}, errors.New("access token not found")
	}

	tokenStr := tokenArr[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
 		}

 		return []byte(config.Get().Jwt.Secret), nil
	})
	if err != nil {
	 return auth.TokenMetadata{}, err
	}

	tokenMetadata, err := extractAccessTokenMetadata(token)
	if err != nil {
		return auth.TokenMetadata{}, err
	}
	return tokenMetadata, nil
}

// AuthMiddleware returns gin.handlerFunc middleware to handle Auth
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if _, err := getAndValidateAccessToken(context.Request); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "invalid token",
			})
			return
		}
		context.Next()
	}
}