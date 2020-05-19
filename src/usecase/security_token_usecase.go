package usecase

import (
	"errors"
	"github.com/google/uuid"
	"root/src/domain/auth"
	"root/src/utils/security"
	"time"
)

// SecurityTokenUseCase implementation of auth.SecurityTokenUseCase
type SecurityTokenUseCase struct {
	SecurityTokenRepo auth.SecurityTokenRepository
}

// GenRefreshToken generates a new refresh token and stores it
func (uc *SecurityTokenUseCase) GenRefreshToken(userID string) (auth.SecurityToken, error) {
	token, err := security.GenRefreshToken(userID)
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
	token, err := security.GenTokenAccessToken(userID)
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

