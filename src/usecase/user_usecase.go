package usecase

import (
	"github.com/google/uuid"
	"sherman/src/domain/auth"
	"sherman/src/utils/exception"
	"sherman/src/utils/security"
	"time"
)

// UserUseCase implementation of auth.UserUseCase
type UserUseCase struct {
	UserRepo auth.UserRepository
}

// NewUserUseCase constructor
func NewUserUseCase(ur auth.UserRepository) auth.UserUseCase {
	var userUseCase auth.UserUseCase = &UserUseCase{
		UserRepo: ur,
	}
	return userUseCase
}

// Register creates a user
func (uc *UserUseCase) Register(user *auth.User) error {
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

// GetUserByID creates a user by id
func (uc *UserUseCase) GetUserByID(id string) (auth.User, error) {
	return uc.UserRepo.GetUserByID(id)
}
