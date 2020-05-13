package presenter

import (
	"root/src/domain"
	"time"
)

// PresentedUser defines struct with public domain.User keys
type PresentedUser struct {
	ID 				uint32		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Active			bool		`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// PresentUser returns a map of public domain.User keys, values
func PresentUser(user *domain.User) PresentedUser {
	return PresentedUser {
		ID:        		user.ID,
		FirstName: 		user.FirstName,
		LastName:  		user.LastName,
		EmailAddress: 	user.EmailAddress,
		Active: 		user.Active,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
	}
}