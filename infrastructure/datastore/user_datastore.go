package datastore

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/smartinsantos/go-auth-api/domain/entity"
	"github.com/smartinsantos/go-auth-api/domain/repository"
	"github.com/smartinsantos/go-auth-api/infrastructure/security"
	"time"
)

type UserDataStore struct {
	db *gorm.DB
}

// UserDataStore implements the repository.UserRepository interface
var _ repository.UserRepository = &UserDataStore{}

// Gets user by id
func (uds *UserDataStore) GetUserById(id uint64) {
	fmt.Println("GetUserById called with id =>", id)
}

// Gets user by email address
func (uds *UserDataStore) GetUserByEmail(email string) {
	fmt.Println("GetUserByEmail called with email =>", email)
}

// Creates a user
func (uds *UserDataStore) CreateUser(user *entity.User) (*entity.User, error) {
	user.ID = uuid.New().ID()
	user.Active = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPassword, err := security.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashPassword)

	err = uds.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Deletes a user
func (uds *UserDataStore) DeleteUser(user *entity.User) {
	fmt.Println("DeleteUser called with =>", user)
}