package datastore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smartinsantos/go-auth-api/domain/entitity"
	"github.com/smartinsantos/go-auth-api/domain/repository"
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
func (uds *UserDataStore) CreateUser(user *entitity.User) (*entitity.User, error) {
	fmt.Println("CreateUser called with =>", user)
	err := uds.db.Create(&user).Error
	if err != nil {
		fmt.Println("ERROR =>", err)
		return nil, err
	}

	return user, nil
}

// Deletes a user
func (uds *UserDataStore) DeleteUser(user *entitity.User) {
	fmt.Println("DeleteUser called with =>", user)
}