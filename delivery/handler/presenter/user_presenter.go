package presenter

import (
	"root/domain"
	"time"
)

type presentedUser struct {
	ID 				uint32		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Active			int			`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// Returns a struct with public domain.User keys
func PresentUser(user *domain.User) *presentedUser {
	return &presentedUser {
		ID:        		user.ID,
		FirstName: 		user.FirstName,
		LastName:  		user.LastName,
		EmailAddress: 	user.EmailAddress,
		Active: 		user.Active,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
	}
}