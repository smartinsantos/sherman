package domain

import (
	"time"
)

// User entity
type User struct {
	ID 				uint32		`gorm:"primary_key; size:36; not null" json:"id"`
	FirstName 		string		`gorm:"size:100;not null;" json:"first_name"`
	LastName 		string		`gorm:"size:100;not null;" json:"last_name"`
	EmailAddress 	string		`gorm:"size:100;not null;unique" json:"email_address"`
	Password 		string		`gorm:"size:100;not null" json:"password"`
	Active			int			`gorm:"not null" json:"active"`
	CreatedAt 		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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