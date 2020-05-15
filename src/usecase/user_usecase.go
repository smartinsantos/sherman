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
func (uc *UserUseCase) Login(user *auth.User) (auth.User, error) {
	record, err := uc.UserRepo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return auth.User{}, err
	}

	err = security.VerifyPassword(record.Password, user.Password)
	if err != nil {
		return record, exception.NewUnAuthorizedError("password doesn't match")
	}

	return record, nil
}
