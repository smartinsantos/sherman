package auth

import (
	"time"
)

// SecurityToken entity struct
type SecurityToken struct {
	ID 				string		`json:"id"`
	UserID 			string		`json:"user_id"`
	Token 			string		`json:"token"`
	Type			string		`json:"type"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

// RefreshTokenType security token type for refresh tokens
const RefreshTokenType = "REFRESH"
// AccessTokenType security token type for access tokens
const AccessTokenType = "ACCESS"
// TokenMetadata struct definition
type TokenMetadata struct {
	UserID 	string
	Type 	string
}

// SecurityTokenRepository interface
type SecurityTokenRepository interface {
	CreateOrUpdateToken(token *SecurityToken) error
	GetTokenByMetadata(tokenMetadata *TokenMetadata) (SecurityToken, error)
}

// SecurityTokenUseCase interface
type SecurityTokenUseCase interface {
	GenRefreshToken(userID string) (SecurityToken, error)
	GenAccessToken(userID string) (SecurityToken, error)
	IsAccessTokenValid(tokenStr string) error
}