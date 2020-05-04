package repository

import "github.com/smartinsantos/go-auth-api/domain/entitity"

type UserRepository interface {
	GetUserById(id uint64)
	GetUserByEmail(email string)
	CreateUser(user *entitity.User)
	DeleteUser(user *entitity.User)
}
