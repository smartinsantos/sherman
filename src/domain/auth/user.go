package auth

import (
	"time"
)

// User entity struct
type User struct {
	ID 				string		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Password 		string		`json:"password"`
	Active			bool		`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// UserRepository interface
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (User, error)
}

// UserUseCase interface
type UserUseCase interface {
	Register(user *User) error
	Login(user *User) (User, string, string, error)
}