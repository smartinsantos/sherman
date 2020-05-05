package entitity

import (
	"time"
)

// User entity
type User struct {
	ID 				uint64		`gorm:"primary_key; size:36; not null" json:"id"`
	FirstName 		string		`gorm:"size:100;not null;" json:"first_name"`
	LastName 		string		`gorm:"size:100;not null;" json:"last_name"`
	EmailAddress 	string		`gorm:"size:100;not null;unique" json:"email_address"`
	Password 		string		`gorm:"size:100;not null" json:"password"`
	Active			int			`gorn:"not null" json:"active"`
	CreatedAt 		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}