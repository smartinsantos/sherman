package security

import (
	"github.com/dgrijalva/jwt-go"
	"root/config"
	"time"
)

func genToken (userID string, exp int64) (string, error) {
	cfg := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat": time.Now().Unix(),
		"exp": exp,
	})

	return token.SignedString([]byte(cfg.Jwt.Secret))
}

// GenTokenAccessToken generates a signed token string
func GenTokenAccessToken(userID string) (string, error) {
	return genToken(userID, time.Now().Add(time.Minute * 15).Unix())
}

// GenRefreshToken generates a signed token string
func GenRefreshToken(userID string) (string, error) {
	return genToken(userID, time.Now().Add(time.Hour * 48).Unix())
}