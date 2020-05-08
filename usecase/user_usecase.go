package usecase

import (
	"github.com/google/uuid"
	"root/app/security"
	"root/domain"
	"time"
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
func (uuc *userUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	// check if the user has been created if not then ...
	user.ID = uuid.New().ID()
	user.Active = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPassword, err := security.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashPassword)

	return uuc.dsUserRepository.CreateUser(user)
}
