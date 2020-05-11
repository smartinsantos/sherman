package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"root/src/app/security"
	"root/src/domain"
)

type userUseCase struct {
	dsUserRepository domain.UserRepository
}

// NewUserUseCase creates a new object representation of domain.UserUseCase interface
func NewUserUseCase(dsUserRepository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		dsUserRepository: dsUserRepository,
	}
}

// Creates a user
func (uuc *userUseCase) CreateUser(user *domain.User) error {
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

	err = uuc.dsUserRepository.CreateUser(user)
	return err
}

// Logs a user in
func (uuc *userUseCase) Login(user *domain.User) (domain.User, error) {
	record, err := uuc.dsUserRepository.GetUserByEmail(user.EmailAddress)
	if err != nil {
		return domain.User{}, err
	}

	err = security.VerifyPassword(record.Password, user.Password)
	if err != nil {
		return record, errors.New("password doesn't match")
	}

	return record, nil
}
