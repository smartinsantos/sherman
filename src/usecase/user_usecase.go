package usecase

import (
	"errors"
	"github.com/google/uuid"
	"root/src/domain/auth"
	"root/src/utils/exception"
	"root/src/utils/security"
	"time"
)

// UserUseCase implementation of auth.UserUseCase
type UserUseCase struct {
	UserRepo auth.UserRepository
	SecurityTokenRepo auth.SecurityTokenRepository
}

// Register creates a user
func (uc *UserUseCase) Register(user *auth.User) error {
	user.ID = uuid.New().String()
	user.Active = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPassword, err := security.Hash(user.Password)
	if  err != nil {
		return err
	}
	user.Password = string(hashPassword)

	err = uc.UserRepo.CreateUser(user)
	return err
}

// Login logs a user in, returns user record and user refresh, access tokens
func (uc *UserUseCase) Login(user *auth.User) (auth.User, string, string, error) {
	userRecord, err := uc.UserRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return auth.User{}, "", "", err
	}

	if err := security.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return auth.User{}, "", "", exception.NewUnAuthorizedError("password doesn't match")
	}

	// refresh token
	grToken, err := security.GenRefreshToken(userRecord.ID)
	if err != nil {
		return auth.User{}, "", "", errors.New("could not generate refresh token")
	}
	refreshToken := auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: userRecord.ID,
		Token: grToken,
		Type: auth.RefreshTokenType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = uc.SecurityTokenRepo.CreateOrUpdateToken(&refreshToken); err != nil {
		return auth.User{}, "", "", errors.New("could not create or update refresh token")
	}

	// access token
	gaToken, err := security.GenTokenAccessToken(userRecord.ID)
	if err != nil {
		return auth.User{}, "", "", errors.New("could not generate access token")
	}
	accessToken := auth.SecurityToken{
		ID: uuid.New().String(),
		UserID: userRecord.ID,
		Token: gaToken,
		Type: auth.AccessTokenType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = uc.SecurityTokenRepo.CreateOrUpdateToken(&accessToken); err != nil {
		return auth.User{}, "", "", errors.New("could not create or update access token")
	}

	return userRecord, refreshToken.Token, accessToken.Token, nil
}
