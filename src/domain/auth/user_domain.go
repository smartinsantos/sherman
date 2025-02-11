package auth

import (
	"time"
)

type (
	// User entity struct
	User struct {
		ID           string    `json:"id"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		EmailAddress string    `json:"email_address"`
		Password     string    `json:"password"`
		Active       bool      `json:"active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	// PresentedUser defines struct with public auth.User keys
	PresentedUser struct {
		ID           string    `json:"id"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		EmailAddress string    `json:"email_address"`
		Active       bool      `json:"active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	// UserRepository interface
	UserRepository interface {
		CreateUser(user *User) error
		GetUserByID(id string) (User, error)
		GetUserByEmail(email string) (User, error)
	}
	// UserUseCase interface
	UserUseCase interface {
		Register(user *User) error
		GetUserByID(id string) (User, error)
		VerifyCredentials(user *User) (User, error)
	}
)
