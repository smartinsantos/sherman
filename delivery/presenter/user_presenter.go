package presenter

import (
	"github.com/smartinsantos/go-auth-api/domain"
	"time"
)

type PresentedUser struct {
	ID 				uint32		`json:"id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	EmailAddress 	string		`json:"email_address"`
	Active			int			`json:"active"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// Exposes Public User keys
func PresentUser(user *domain.User) *PresentedUser {
	return &PresentedUser {
		ID:        		user.ID,
		FirstName: 		user.FirstName,
		LastName:  		user.LastName,
		EmailAddress: 	user.EmailAddress,
		Active: 		user.Active,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
	}
}