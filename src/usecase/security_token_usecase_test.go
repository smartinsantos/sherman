package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sherman/mocks"
	_ "sherman/src/app/testing"
	"sherman/src/domain/auth"
	"testing"
	"time"
)

type securityTokenUseCaseMockDeps struct {
	securityTokenRepository *mocks.SecurityTokenRepository
	securityService         *mocks.Security
}

func genSecurityTokenUseCase() (auth.SecurityTokenUseCase, securityTokenUseCaseMockDeps) {
	stucDeps := securityTokenUseCaseMockDeps{
		securityTokenRepository: new(mocks.SecurityTokenRepository),
		securityService:         new(mocks.Security),
	}

	stuc := NewSecurityTokenUseCase(
		stucDeps.securityTokenRepository,
		stucDeps.securityService,
	)

	return stuc, stucDeps
}

func TestGenRefreshToken(t *testing.T) {
	mockUserID := "some-user-id"
	mockToken := "some-token"

	t.Run("it should succeed", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()
		stucDeps.securityTokenRepository.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		stucDeps.securityService.
			On(
				"GenToken",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int64"),
				mock.AnythingOfType("int64"),
			).
			Return(mockToken, nil)

		refreshToken, err := stuc.GenRefreshToken(mockUserID)

		assert.NoError(t, err)
		assert.NotEmpty(t, refreshToken.ID)
		assert.EqualValues(t, mockUserID, refreshToken.UserID)
		assert.EqualValues(t, auth.RefreshTokenType, refreshToken.Type)
		assert.NotEmpty(t, refreshToken.CreatedAt)
		assert.NotEmpty(t, refreshToken.UpdatedAt)
	})

	t.Run("it should return an error", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()
		stucDeps.securityTokenRepository.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		mockError := errors.New("some error")
		stucDeps.securityService.
			On(
				"GenToken",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int64"),
				mock.AnythingOfType("int64"),
			).
			Return("", mockError)

		_, err := stuc.GenRefreshToken(mockUserID)

		if assert.Error(t, err) {
			assert.Equal(t, "could not generate refresh token", err.Error())
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()
		mockError := errors.New("some error")
		stucDeps.securityTokenRepository.On("CreateOrUpdateToken", mock.Anything).Return(mockError)
		stucDeps.securityService.
			On(
				"GenToken",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int64"),
				mock.AnythingOfType("int64"),
			).
			Return(mockToken, nil)

		_, err := stuc.GenRefreshToken(mockUserID)

		if assert.Error(t, err) {
			assert.Equal(t, "could not create or update refresh token", err.Error())
		}
	})
}

func TestGenAccessToken(t *testing.T) {
	mockUserID := "some-user-id"
	mockToken := "some-token"

	t.Run("it should succeed", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()
		stucDeps.securityTokenRepository.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		stucDeps.securityService.
			On(
				"GenToken",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int64"),
				mock.AnythingOfType("int64"),
			).
			Return(mockToken, nil)

		refreshToken, err := stuc.GenAccessToken(mockUserID)

		assert.NoError(t, err)
		assert.NotEmpty(t, refreshToken.ID)
		assert.EqualValues(t, mockUserID, refreshToken.UserID)
		assert.EqualValues(t, auth.AccessTokenType, refreshToken.Type)
		assert.NotEmpty(t, refreshToken.CreatedAt)
		assert.NotEmpty(t, refreshToken.UpdatedAt)
	})

	t.Run("it should return an error", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()

		stucDeps.securityTokenRepository.On("CreateOrUpdateToken", mock.Anything).Return(nil)
		mockError := errors.New("some error")
		stucDeps.securityService.
			On(
				"GenToken",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int64"),
				mock.AnythingOfType("int64"),
			).
			Return("", mockError)

		_, err := stuc.GenAccessToken(mockUserID)

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
		stuc, stucDeps := genSecurityTokenUseCase()
		stucDeps.securityTokenRepository.On("GetTokenByMetadata", mock.Anything).Return(mockSecurityToken, nil)

		tokenStored := stuc.IsRefreshTokenStored(mockRefreshTokenMetaData)

		assert.EqualValues(t, true, tokenStored)
	})

	t.Run("it should return an error", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()
		stucDeps.securityTokenRepository.
			On("GetTokenByMetadata", mock.Anything).
			Return(auth.SecurityToken{}, errors.New("some error"))

		tokenStored := stuc.IsRefreshTokenStored(mockRefreshTokenMetaData)

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
		stuc, stucDeps := genSecurityTokenUseCase()
		stucDeps.securityTokenRepository.On("RemoveTokenByMetadata", mock.Anything).Return(nil)

		err := stuc.RemoveRefreshToken(mockRefreshTokenMetaData)

		assert.NoError(t, err)
	})

	t.Run("it should return an error", func(t *testing.T) {
		stuc, stucDeps := genSecurityTokenUseCase()
		mockError := errors.New("some error")
		stucDeps.securityTokenRepository.On("RemoveTokenByMetadata", mock.Anything).Return(mockError)

		err := stuc.RemoveRefreshToken(mockRefreshTokenMetaData)

		if assert.Error(t, err) {
			assert.EqualValues(t, mockError, err)
		}
	})
}
