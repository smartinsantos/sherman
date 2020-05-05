package entity

import (
	"time"
	"strings"
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

type Users []User

type PresentedUser  struct {
	ID 				uint32		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Active			int			`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// Exposes Public User keys
func (u *User) Presenter() *PresentedUser {
	return &PresentedUser {
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		EmailAddress: u.EmailAddress,
		Active: u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
		default:
			if u.FirstName == "" {
				errorMessages["first_name_required"] = "First name is required"
			}
			if u.LastName == "" {
				errorMessages["last_name_required"] = "Last name is required"
			}
			if u.Password == "" {
				errorMessages["password_required"] = "Password is required"
			}
			if u.EmailAddress == "" {
				errorMessages["email_address_required"] = "Email Address is required"
			}
	}
	return errorMessages
}

