package auth

import (
	"time"
)

const (
	// RefreshTokenType constant security token type for refresh tokens
	RefreshTokenType = "REFRESH"
	// AccessTokenType constant security token type for access tokens
	AccessTokenType = "ACCESS"
)

type (
	// SecurityToken entity struct
	SecurityToken struct {
		ID 				string		`json:"id"`
		UserID 			string		`json:"user_id"`
		Token 			string		`json:"token"`
		Type			string		`json:"type"`
		CreatedAt 		time.Time	`json:"created_at"`
		UpdatedAt 		time.Time	`json:"updated_at"`
	}
	// TokenMetadata struct definition
	TokenMetadata struct {
		UserID 	string
		Type 	string
		Token	string
	}
	// SecurityTokenRepository interface
	SecurityTokenRepository interface {
		CreateOrUpdateToken(token *SecurityToken) error
		GetTokenByMetadata(tokenMetadata *TokenMetadata) (SecurityToken, error)
	}
	// SecurityTokenUseCase interface
	SecurityTokenUseCase interface {
		GenRefreshToken(userID string) (SecurityToken, error)
		GenAccessToken(userID string) (SecurityToken, error)
		IsRefreshTokenStored(refreshTokenMetadata *TokenMetadata) bool
	}
)


