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
	UserRepo auth.UserRepository
}

// NewUserUseCase constructor
func NewUserUseCase(ur auth.UserRepository) auth.UserUseCase {
	return &userUseCase{
		UserRepo: ur,
	}
}

// Register creates a user
func (uc *userUseCase) Register(user *auth.User) error {
	user.ID = uuid.New().String()
	user.Active = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPassword, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	return uc.UserRepo.CreateUser(user)
}

// VerifyCredentials verifies a user credentials
func (uc *userUseCase) VerifyCredentials(user *auth.User) (auth.User, error) {
	userRecord, err := uc.UserRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return auth.User{}, err
	}

	if err := security.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return auth.User{}, exception.NewUnAuthorizedError("password doesn't match")
	}

	return userRecord, nil
}

// GetUserByID creates a user by id
func (uc *userUseCase) GetUserByID(id string) (auth.User, error) {
	return uc.UserRepo.GetUserByID(id)
}
