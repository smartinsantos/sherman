package auth

import (
	"time"
)

// SecurityToken entity struct
type SecurityToken struct {
	ID 				uint32		`json:"id"`
	UserID 			string		`json:"user_id"`
	Token 			string		`json:"token"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// SecurityTokenRepository interface
type SecurityTokenRepository interface {
	CreateToken(token *SecurityToken) error
	//UpdateToken(token *SecurityToken) error
	//GetTokenByUserId(userId string) (SecurityToken, error)
}
