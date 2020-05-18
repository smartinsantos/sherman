package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"root/config"
	"strings"
	"time"
)

func genToken (userID string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat": time.Now().Unix(),
		"exp": exp,
	})

	return token.SignedString([]byte(config.Get().Jwt.Secret))
}

// GenTokenAccessToken generates a signed token string
func GenTokenAccessToken(userID string) (string, error) {
	return genToken(userID, time.Now().Add(time.Minute * 15).Unix())
}

//@TODO WORKING ON THIS
// ExtractAccessToken extracts a token from http request
func ExtractAccessToken(req *http.Request) string {
  bearToken := req.Header.Get("Authorization")
  tokenArr := strings.Split(bearToken, " ")
  if len(tokenArr) != 2 {
  	return ""
  }
  return tokenArr[1]
}

// VerifyAccessTokenSignature Verifys if an access token has the right signature
func VerifyAccessTokenSignature(req *http.Request) (*jwt.Token, error) {
	tokenStr := ExtractAccessToken(req)
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

// GenRefreshToken generates a signed token string
func GenRefreshToken(userID string) (string, error) {
	return genToken(userID, time.Now().Add(time.Hour * 48).Unix())
}
