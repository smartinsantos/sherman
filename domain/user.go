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
	Active			int			`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// User use case interface
type UserUseCase interface {
	//GetUserById(id uint64) (*User, error)
	//GetUserByEmail(email string) (*User, error)
	//UpdateUser(user *User) (*User, error)
	CreateUser(user *User) (*User, error)
	//DeleteUser(user *User) error
}

// User repository interface
type UserRepository interface {
	//GetUserById(id uint64) (*User, error)
	//GetUserByEmail(email string) (*User, error)
	//UpdateUser(user *User) (*User, error)
	CreateUser(user *User) (*User, error)
	//DeleteUser(user *User) error
}