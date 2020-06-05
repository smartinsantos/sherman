package usecase

import (
	"github.com/google/uuid"
	"sherman/src/app/utils/exception"
	"sherman/src/app/utils/security"
	"sherman/src/domain/auth"
	"time"
)

// UserUseCase implementation of auth.UserUseCase
type userUseCase struct {
	userRepo     auth.UserRepository
	passwordUtil security.PasswordUtil
}

// NewUserUseCase constructor
func NewUserUseCase(ur auth.UserRepository, p security.PasswordUtil) auth.UserUseCase {
	return &userUseCase{
		userRepo:     ur,
		passwordUtil: p,
	}
}

// Register creates a user
func (uc *userUseCase) Register(user *auth.User) error {
	user.ID = uuid.New().String()
	user.Active = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPassword, err := uc.passwordUtil.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	return uc.userRepo.CreateUser(user)
}

// VerifyCredentials verifies a user credentials
func (uc *userUseCase) VerifyCredentials(user *auth.User) (auth.User, error) {
	userRecord, err := uc.userRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return auth.User{}, err
	}

	if err := uc.passwordUtil.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return auth.User{}, exception.NewUnAuthorizedError("password doesn't match")
	}

	return userRecord, nil
}

// GetUserByID creates a user by id
func (uc *userUseCase) GetUserByID(id string) (auth.User, error) {
	return uc.userRepo.GetUserByID(id)
}
