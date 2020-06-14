package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"sherman/src/app/config"
	"sherman/src/domain/auth"
	"strings"
)

func (s *service) extractTokenMetadata(token *jwt.Token) (auth.TokenMetadata, error) {
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
		Type:   tokenType,
		Token:  token.Raw,
	}, nil
}

func (s *service) parseTokenString(ts string) (*jwt.Token, error) {
	return jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(config.Get().Jwt.Secret), nil
	})
}

// GenToken generates a jwt.token
func (s *service) GenToken(userID, tokenType string, iat, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"iat":     iat,
		"exp":     exp,
	})

	return token.SignedString([]byte(config.Get().Jwt.Secret))
}

// GetAndValidateAccessToken gets the access token from echo.Context and verifies its signature
func (s *service) GetAndValidateAccessToken(ctx echo.Context) (auth.TokenMetadata, error) {
	tokenHeader := ctx.Request().Header.Get("Authorization")
	tokenHeaderArr := strings.Split(tokenHeader, " ")
	if len(tokenHeaderArr) != 2 {
		return auth.TokenMetadata{}, errors.New("access token not found")
	}
	token, err := s.parseTokenString(tokenHeaderArr[1])
	if err != nil {
		return auth.TokenMetadata{}, errors.New("invalid token")
	}

	tokenMetadata, err := s.extractTokenMetadata(token)
	if err != nil {
		return auth.TokenMetadata{}, errors.New("invalid token")
	}
	return tokenMetadata, nil
}

// GetAndValidateRefreshToken gets the refresh token from echo.Context and verifies its signature
func (s *service) GetAndValidateRefreshToken(ctx echo.Context) (auth.TokenMetadata, error) {
	refreshTokenCookie, err := ctx.Request().Cookie("REFRESH_TOKEN")
	if err != nil {
		return auth.TokenMetadata{}, errors.New("refresh token not found")
	}

	token, err := s.parseTokenString(refreshTokenCookie.Value)
	if err != nil {
		return auth.TokenMetadata{}, errors.New("invalid refresh token")
	}

	tokenMetadata, err := s.extractTokenMetadata(token)
	if err != nil {
		return auth.TokenMetadata{}, errors.New("invalid refresh token")
	}
	return tokenMetadata, nil
}
