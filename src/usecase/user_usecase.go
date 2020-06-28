package usecase

import (
	"github.com/google/uuid"
	"sherman/src/app/utils/terr"
	"sherman/src/domain/auth"
	"sherman/src/service/security"
	"time"
)

// UserUseCase implementation of auth.UserUseCase
type userUseCase struct {
	userRepo auth.UserRepository
	security security.Security
}

// NewUserUseCase constructor
func NewUserUseCase(ur auth.UserRepository, ss security.Security) auth.UserUseCase {
	return &userUseCase{
		userRepo: ur,
		security: ss,
	}
}

// Register creates a user
func (uc *userUseCase) Register(user *auth.User) error {
	user.ID = uuid.New().String()
	user.Active = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPassword, err := uc.security.Hash(user.Password)
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

	if err := uc.security.VerifyPassword(userRecord.Password, user.Password); err != nil {
		return auth.User{}, terr.NewUnAuthorizedError("password doesn't match")
	}

	return userRecord, nil
}

// GetUserByID creates a user by id
func (uc *userUseCase) GetUserByID(id string) (auth.User, error) {
	return uc.userRepo.GetUserByID(id)
}
