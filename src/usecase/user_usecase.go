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
	user.ID = uuid.New().ID()
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

// Login logs a user in, returns user record and user token[TODO]
func (uc *UserUseCase) Login(user *auth.User) (auth.User, string, error) {
	userRecord, err := uc.UserRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return auth.User{}, auth.SecurityToken{}.Token, err
	}

	if err := security.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return auth.User{}, auth.SecurityToken{}.Token, exception.NewUnAuthorizedError("password doesn't match")
	}

	gToken, err := security.GenToken(userRecord.ID)
	if err != nil {
		return auth.User{}, auth.SecurityToken{}.Token, errors.New("could not generate token")
	}

	securityToken := auth.SecurityToken{
		ID: uuid.New().ID(),
		UserID: userRecord.ID,
		Token: gToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = uc.SecurityTokenRepo.CreateOrUpdateToken(&securityToken); err != nil {
		return auth.User{}, auth.SecurityToken{}.Token, errors.New("could not create or update token")
	}

	return userRecord, gToken, nil
}
