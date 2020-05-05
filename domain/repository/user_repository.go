package repository

import (
	"github.com/smartinsantos/go-auth-api/domain/entitity"
)

// User Repository interface
type UserRepository interface {
	GetUserById(id uint64)
	GetUserByEmail(email string)
	CreateUser(user *entitity.User) (*entitity.User, error)
	DeleteUser(user *entitity.User)
}