package repository

import (
	"github.com/smartinsantos/go-auth-api/domain/entity"
)

// User Repository interface
type UserRepository interface {
	GetUserById(id uint64)
	GetUserByEmail(email string)
	CreateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User)
}