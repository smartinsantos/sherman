package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"root/config"
	"root/src/domain/auth"
	"strings"
	"time"
)

// GenTokenAccessToken generates a signed token string
func GenTokenAccessToken(userID string) (string, error) {
	return genToken(userID, time.Now().Add(time.Minute * 15).Unix())
}

// ExtractAccessTokenMetadata extracts metadata of access token from *http.Request
func ExtractAccessTokenMetadata(r *http.Request) (auth.TokenMetadata, error) {
	token, err := getAccessTokenFromRequest(r)
	if err != nil {
		return auth.TokenMetadata{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return auth.TokenMetadata{}, err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return auth.TokenMetadata{}, err
	}

	return auth.TokenMetadata{
		UserID: userID,
		Type: auth.AccessTokenType,
	}, nil
}

// GenRefreshToken generates a signed token string
func GenRefreshToken(userID string) (string, error) {
	return genToken(userID, time.Now().Add(time.Hour * 48).Unix())
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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
 		}

 		return []byte(config.Get().Jwt.Secret), nil
	})

	if err != nil {
	 return nil, err
	}

	return token, nil
}

// genToken generates a jwt.token
func genToken(userID string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat": time.Now().Unix(),
		"exp": exp,
	})

	return token.SignedString([]byte(config.Get().Jwt.Secret))
}