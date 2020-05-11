package domain

import (
	"time"
)

// User entity
type User struct {
	ID 				uint32		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Password 		string		`json:"password"`
	Active			bool		`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// User use case interface
type UserUseCase interface {
	CreateUser(user *User) error
	Login(user *User) (User, error)
}

// User repository interface
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (User, error)
}