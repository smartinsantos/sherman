package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"root/config"
	"root/src/domain/auth"
	"strings"
	"time"
)

// GenToken generates a jwt.token
func GenToken(userID string, tokenType string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"type": tokenType,
		"iat": time.Now().Unix(),
		"exp": exp,
	})

	return token.SignedString([]byte(config.Get().Jwt.Secret))
}

// GetAndValidateAccessToken gets the access token from *gin.Context and verifies its signature
func GetAndValidateAccessToken(ctx *gin.Context) (auth.TokenMetadata, error) {
	bearToken := ctx.Request.Header.Get("Authorization")
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

	tokenMetadata, err := extractTokenMetadata(token)
	if err != nil {
		return auth.TokenMetadata{}, err
	}
	return tokenMetadata, nil
}

// GetAndValidateRefreshToken gets the refresh token from *gin.Context and verifies its signature
func GetAndValidateRefreshToken(ctx *gin.Context) (auth.TokenMetadata, error) {
	refreshTokenCookie, err := ctx.Request.Cookie("REFRESH_TOKEN")
	if err != nil {
		return auth.TokenMetadata{}, errors.New("refresh token not found")
	}

	tokenStr := refreshTokenCookie.Value
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.Get().Jwt.Secret), nil
	})
	if err != nil {
	 return auth.TokenMetadata{}, err
	}

	tokenMetadata, err := extractTokenMetadata(token)
	if err != nil {
		return auth.TokenMetadata{}, err
	}
	return tokenMetadata, nil
}

// extractTokenMetadata  extracts metadata of *jwt.Token
func extractTokenMetadata(token *jwt.Token) (auth.TokenMetadata, error) {
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
		Token: token.Raw,
	}, nil
}