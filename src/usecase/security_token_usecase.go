package usecase

import (
	"errors"
	"github.com/google/uuid"
	"sherman/src/domain/auth"
	"sherman/src/utils/security"
	"time"
)

// SecurityTokenUseCase implementation of auth.SecurityTokenUseCase
type SecurityTokenUseCase struct {
	SecurityTokenRepo auth.SecurityTokenRepository
}

// GenRefreshToken generates a new refresh token and stores it
func (uc *SecurityTokenUseCase) GenRefreshToken(userID string) (auth.SecurityToken, error) {
	duration := time.Hour * time.Duration(48)
	token, err := security.GenToken(userID, auth.RefreshTokenType, time.Now().Add(duration).Unix())
	if err != nil {
		return auth.SecurityToken{}, errors.New("could not generate refresh token")
	}

	refreshToken := auth.SecurityToken{
		ID:        uuid.New().String(),
		UserID:    userID,
		Token:     token,
		Type:      auth.RefreshTokenType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = uc.SecurityTokenRepo.CreateOrUpdateToken(&refreshToken); err != nil {
		return auth.SecurityToken{}, errors.New("could not create or update refresh token")
	}

	return refreshToken, nil
}

// GenAccessToken generates a new access token
func (uc *SecurityTokenUseCase) GenAccessToken(userID string) (auth.SecurityToken, error) {
	duration := time.Minute * time.Duration(15)
	token, err := security.GenToken(userID, auth.AccessTokenType, time.Now().Add(duration).Unix())
	if err != nil {
		return auth.SecurityToken{}, errors.New("could not generate access token")
	}

	accessToken := auth.SecurityToken{
		ID:        uuid.New().String(),
		UserID:    userID,
		Token:     token,
		Type:      auth.AccessTokenType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return accessToken, nil
}

// IsRefreshTokenStored checks if a refresh token is persisted in the datastore
func (uc *SecurityTokenUseCase) IsRefreshTokenStored(refreshTokenMetadata *auth.TokenMetadata) bool {
	_, err := uc.SecurityTokenRepo.GetTokenByMetadata(refreshTokenMetadata)
	return err == nil
}

// RemoveRefreshToken removes a refresh token from the datastore
func (uc *SecurityTokenUseCase) RemoveRefreshToken(refreshTokenMetadata *auth.TokenMetadata) error {
	return uc.SecurityTokenRepo.RemoveTokenByMetadata(refreshTokenMetadata)
}
