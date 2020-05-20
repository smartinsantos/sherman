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

// VerifyCredentials verifies user credentials
func (uc *UserUseCase) VerifyCredentials(user *auth.User) (auth.User, error) {
	userRecord, err := uc.UserRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return auth.User{}, err
	}

	if err := security.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return auth.User{}, exception.NewUnAuthorizedError("password doesn't match")
	}

	return userRecord, nil
}