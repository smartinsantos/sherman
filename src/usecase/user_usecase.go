package usecase

import (
	"github.com/google/uuid"
	"root/src/domain/auth"
	"root/src/utils/exception"
	"root/src/utils/security"
	"time"
)

// UserUseCase implementation of auth.UserUseCase
type UserUseCase struct {
	UserRepo auth.UserRepository
	SecurityTokenUseCase auth.SecurityTokenUseCase
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
func (uc *UserUseCase) Login(user *auth.User) (string, string, error) {
	userRecord, err := uc.UserRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return "", "", err
	}

	if err := security.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return "", "", exception.NewUnAuthorizedError("password doesn't match")
	}

	refreshToken, err := uc.SecurityTokenUseCase.GenRefreshToken(userRecord.ID)
	if err != nil {
		return "", "", err
	}

	accessToken, err := uc.SecurityTokenUseCase.GenAccessToken(userRecord.ID)
	if err != nil {
		return "", "", err
	}

	return refreshToken.Token, accessToken.Token, nil
}