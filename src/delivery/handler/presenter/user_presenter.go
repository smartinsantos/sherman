package presenter

import (
	"root/src/domain/auth"
	"time"
)

// PresentedUser defines struct with public auth.User keys
type PresentedUser struct {
	ID 				string		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Active			bool		`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// PresentUser returns a map of public auth.User keys, values
func PresentUser(user *auth.User) PresentedUser {
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