package usecase

import (
	"errors"
	"github.com/google/uuid"
	"sherman/src/app/utils/security"
	"sherman/src/domain/auth"
	"time"
)

// SecurityTokenUseCase implementation of auth.SecurityTokenUseCase
type securityTokenUseCase struct {
	securityTokenRepo auth.SecurityTokenRepository
	tokenUtil         security.TokenUtil
}

// NewSecurityTokenUseCase constructor
func NewSecurityTokenUseCase(str auth.SecurityTokenRepository, tu security.TokenUtil) auth.SecurityTokenUseCase {
	return &securityTokenUseCase{
		securityTokenRepo: str,
		tokenUtil:         tu,
	}
}

// GenRefreshToken generates a new refresh token and stores it
func (uc *securityTokenUseCase) GenRefreshToken(userID string) (auth.SecurityToken, error) {
	duration := time.Hour * time.Duration(48)
	token, err := uc.tokenUtil.GenToken(userID, auth.RefreshTokenType, time.Now().Add(duration).Unix())
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
	if err = uc.securityTokenRepo.CreateOrUpdateToken(&refreshToken); err != nil {
		return auth.SecurityToken{}, errors.New("could not create or update refresh token")
	}

	return refreshToken, nil
}

// GenAccessToken generates a new access token
func (uc *securityTokenUseCase) GenAccessToken(userID string) (auth.SecurityToken, error) {
	duration := time.Minute * time.Duration(15)
	token, err := uc.tokenUtil.GenToken(userID, auth.AccessTokenType, time.Now().Add(duration).Unix())
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
func (uc *securityTokenUseCase) IsRefreshTokenStored(refreshTokenMetadata *auth.TokenMetadata) bool {
	_, err := uc.securityTokenRepo.GetTokenByMetadata(refreshTokenMetadata)
	return err == nil
}

// RemoveRefreshToken removes a refresh token from the datastore
func (uc *securityTokenUseCase) RemoveRefreshToken(refreshTokenMetadata *auth.TokenMetadata) error {
	return uc.securityTokenRepo.RemoveTokenByMetadata(refreshTokenMetadata)
}
