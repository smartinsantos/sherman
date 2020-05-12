package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"root/src/app/security"
	"root/src/domain"
)

// Implementation of domain.UserUseCase
type UserUseCase struct {
	Repo domain.UserRepository
}

// Creates a user
func (uc *UserUseCase) CreateUser(user *domain.User) error {
	var err error

	user.ID = uuid.New().ID()
	user.Active = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if hashPassword, err := security.Hash(user.Password); err != nil {
		return err
	} else {
		user.Password = string(hashPassword)
	}

	err = uc.Repo.CreateUser(user)
	return err
}

// Logs a user in
func (uc *UserUseCase) Login(user *domain.User) (domain.User, error) {
	record, err := uc.Repo.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return domain.User{}, err
	}

	err = security.VerifyPassword(record.Password, user.Password)
	if err != nil {
		return record, errors.New("password doesn't match")
	}

	return record, nil
}
