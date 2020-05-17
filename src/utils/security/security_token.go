package security

import (
	"github.com/dgrijalva/jwt-go"
	"root/config"
	"time"
)

// GenToken generates a signed token string
func GenToken(userID uint32) (string, error) {
	cfg := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	return token.SignedString([]byte(cfg.Jwt.Secret))
}