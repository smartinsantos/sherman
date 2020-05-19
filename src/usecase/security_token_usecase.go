package usecase

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"root/config"
	"root/src/domain/auth"
	"time"
)

// SecurityTokenUseCase implementation of auth.SecurityTokenUseCase
type SecurityTokenUseCase struct {
	SecurityTokenRepo auth.SecurityTokenRepository
}

// GenRefreshToken generates a new refresh token and stores it
func (uc *SecurityTokenUseCase) GenRefreshToken(userID string) (auth.SecurityToken, error) {
	token, err := genToken(userID, time.Now().Add(time.Hour * 48).Unix())
	if err != nil {
		return auth.SecurityToken{}, errors.New("could not generate refresh token")
	}

	refreshToken := auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: userID,
		Token: token,
		Type: auth.RefreshTokenType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = uc.SecurityTokenRepo.CreateOrUpdateToken(&refreshToken); err != nil {
		return auth.SecurityToken{}, errors.New("could not create or update refresh token")
	}

	return refreshToken, nil
}

// GenAccessToken generates a new access token and stores it
func (uc *SecurityTokenUseCase) GenAccessToken(userID string) (auth.SecurityToken, error) {
	token, err := genToken(userID, time.Now().Add(time.Minute * 15).Unix())
	if err != nil {
		return auth.SecurityToken{}, errors.New("could not generate access token")
	}

	accessToken := auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: userID,
		Token: token,
		Type: auth.AccessTokenType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = uc.SecurityTokenRepo.CreateOrUpdateToken(&accessToken); err != nil {
		return auth.SecurityToken{}, errors.New("could not create or update access token")
	}

	return accessToken, nil
}

func (uc *SecurityTokenUseCase) IsAccessTokenValid(tokenStr string) error {
	//@TODO: implement
	return nil
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