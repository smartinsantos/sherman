package datastore

import (
	"fmt"
	"github.com/smartinsantos/go-auth-api/domain/entitity"
	"github.com/smartinsantos/go-auth-api/domain/repository"
)

type UserDataStore struct {}

// UserDataStore implements the repository.UserRepository interface
var _ repository.UserRepository = &UserDataStore{}

func (uds *UserDataStore) GetUserById(id uint64) {
	fmt.Println("GetUserById called with id =>", id)
}

func (uds *UserDataStore) GetUserByEmail(email string) {
	fmt.Println("GetUserByEmail called with email =>", email)
}

func (uds *UserDataStore) CreateUser(user *entitity.User) {
	fmt.Println("CreateUser called with =>", user)
}

func (uds *UserDataStore) DeleteUser(user *entitity.User) {
	fmt.Println("DeleteUser called with =>", user)
}