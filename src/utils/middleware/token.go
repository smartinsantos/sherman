package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"root/config"
	"root/src/app/registry"
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

	return auth.TokenMetadata{
		UserID: userID,
		Type: auth.AccessTokenType,
	}, nil
}

// getAccessTokenFromRequest verifies signature and extracts the token from an http request
func getAccessTokenFromRequest(req *http.Request) (*jwt.Token, error) {
	bearToken := req.Header.Get("Authorization")
	tokenArr := strings.Split(bearToken, " ")
	if len(tokenArr) != 2 {
		return nil, errors.New("access token not found")
	}

	tokenStr := tokenArr[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
 		}

 		return []byte(config.Get().Jwt.Secret), nil
	})

	if err != nil {
	 return nil, err
	}

	return token, nil
}

// AuthMiddleware returns gin.handlerFunc middleware to handle Auth
func AuthMiddleware() gin.HandlerFunc {
	diContainer, _ := registry.GetAppContainer()

	return func(context *gin.Context) {
		if diContainer == nil {
			context.AbortWithError(http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		token, err := getAccessTokenFromRequest(context.Request)
		if err != nil {
			log.Println(err)
			context.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		fmt.Println("TOKEN", token)
		tokenMetadata, err := extractAccessTokenMetadata(token)
		if err != nil {
			log.Println(err)
			context.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		fmt.Println("tokenMetadata =>", tokenMetadata)
		//securityTokenUseCase := diContainer.Get("security-token-usecase").(*usecase.SecurityTokenUseCase)

		context.Next()
	}
}