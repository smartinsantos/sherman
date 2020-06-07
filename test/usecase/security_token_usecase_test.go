package test_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sherman/mocks"
	"sherman/src/domain/auth"
	"sherman/src/usecase"
	"testing"
	"time"
)

func TestGenRefreshToken(t *testing.T) {
	mockUserId := "some-user-id"
	mockToken := "some-token"

	t.Run("it should succeed", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		mockSecurityService := new(mocks.Security)
		mockSecurityService.
		On(
			"GenToken",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).
		Return(mockToken, nil)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		refreshToken, err := securityTokenUseCase.GenRefreshToken(mockUserId)

		assert.NoError(t, err)
		assert.NotEmpty(t, refreshToken.ID)
		assert.EqualValues(t, mockUserId, refreshToken.UserID)
		assert.EqualValues(t, auth.RefreshTokenType, refreshToken.Type)
		assert.NotEmpty(t, refreshToken.CreatedAt)
		assert.NotEmpty(t, refreshToken.UpdatedAt)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		mockSecurityService := new(mocks.Security)
		mockError := errors.New("some error")
		mockSecurityService.
		On(
			"GenToken",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).
		Return("", mockError)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		_, err := securityTokenUseCase.GenRefreshToken(mockUserId)

		if assert.Error(t, err) {
			assert.Equal(t, "could not generate refresh token", err.Error())
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockError := errors.New("some error")
		mockSecurityTokenRepo.On("CreateOrUpdateToken", mock.Anything).Return(mockError)
		mockSecurityService := new(mocks.Security)
		mockSecurityService.
		On(
			"GenToken",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).
		Return(mockToken, nil)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		_, err := securityTokenUseCase.GenRefreshToken(mockUserId)

		if assert.Error(t, err) {
			assert.Equal(t, "could not create or update refresh token", err.Error())
		}
	})
}

func TestGenAccessToken(t *testing.T) {
	mockUserId := "some-user-id"
	mockToken := "some-token"

	t.Run("it should succeed", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		mockSecurityService := new(mocks.Security)
		mockSecurityService.
		On(
			"GenToken",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).
		Return(mockToken, nil)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		refreshToken, err := securityTokenUseCase.GenAccessToken(mockUserId)

		assert.NoError(t, err)
		assert.NotEmpty(t, refreshToken.ID)
		assert.EqualValues(t, mockUserId, refreshToken.UserID)
		assert.EqualValues(t, auth.AccessTokenType, refreshToken.Type)
		assert.NotEmpty(t, refreshToken.CreatedAt)
		assert.NotEmpty(t, refreshToken.UpdatedAt)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		mockSecurityService := new(mocks.Security)
		mockError := errors.New("some error")
		mockSecurityService.
		On(
			"GenToken",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).
		Return("", mockError)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		_, err := securityTokenUseCase.GenAccessToken(mockUserId)

		if assert.Error(t, err) {
			assert.Equal(t, "could not generate access token", err.Error())
		}
	})
}

func TestIsRefreshTokenStored(t *testing.T) {
	mockRefreshTokenMetaData := &auth.TokenMetadata{
		UserID: "some-user-id",
		Type:   auth.RefreshTokenType,
		Token:  "some-token",
	}

	now := time.Now()
	mockSecurityToken := auth.SecurityToken{
		ID:        "some-id",
		UserID:    "some-user-id",
		Token:     "some-token",
		Type:      auth.RefreshTokenType,
		CreatedAt: now,
		UpdatedAt: now,
	}

	t.Run("it should succeed", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.On("GetTokenByMetadata", mock.Anything).Return(mockSecurityToken, nil)
		mockSecurityService := new(mocks.Security)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		tokenStored := securityTokenUseCase.IsRefreshTokenStored(mockRefreshTokenMetaData)

		assert.EqualValues(t, true, tokenStored)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.
			On("GetTokenByMetadata", mock.Anything).
			Return(auth.SecurityToken{}, errors.New("some error"))
		mockSecurityService := new(mocks.Security)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		tokenStored := securityTokenUseCase.IsRefreshTokenStored(mockRefreshTokenMetaData)

		assert.EqualValues(t, false, tokenStored)
	})
}

func TestRemoveRefreshToken(t *testing.T) {
	mockRefreshTokenMetaData := &auth.TokenMetadata{
		UserID: "some-user-id",
		Type:   auth.RefreshTokenType,
		Token:  "some-token",
	}

	t.Run("it should succeed", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockSecurityTokenRepo.On("RemoveTokenByMetadata", mock.Anything).Return(nil)
		mockSecurityService := new(mocks.Security)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		err := securityTokenUseCase.RemoveRefreshToken(mockRefreshTokenMetaData)

		assert.NoError(t, err)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockSecurityTokenRepo := new(mocks.SecurityTokenRepository)
		mockError := errors.New("some error")
		mockSecurityTokenRepo.On("RemoveTokenByMetadata", mock.Anything).Return(mockError)
		mockSecurityService := new(mocks.Security)

		securityTokenUseCase := usecase.NewSecurityTokenUseCase(mockSecurityTokenRepo, mockSecurityService)

		err := securityTokenUseCase.RemoveRefreshToken(mockRefreshTokenMetaData)

		if assert.Error(t, err) {
			assert.EqualValues(t, mockError, err)
		}
	})
}