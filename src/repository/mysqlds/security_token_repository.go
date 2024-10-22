package mysqlds

import (
	"database/sql"
	"sherman/src/app/utils/terr"
	"sherman/src/domain/auth"
)

// securityTokenRepository sql implementation of auth.SecurityTokenRepository
type securityTokenRepository struct {
	DB *sql.DB
}

// NewSecurityTokenRepository constructor
func NewSecurityTokenRepository(db *sql.DB) auth.SecurityTokenRepository {
	return &securityTokenRepository{
		DB: db,
	}
}

// CreateOrUpdateToken persist a auth.SecurityToken in the datastore
func (r *securityTokenRepository) CreateOrUpdateToken(token *auth.SecurityToken) error {
	var err error
	var query string
	var existingToken auth.SecurityToken

	// find token id if it exist
	query = `SELECT id FROM security_tokens WHERE user_id = ? AND type = ? LIMIT 1`
	row := r.DB.QueryRow(query, token.UserID, token.Type)
	_ = row.Scan(&existingToken.ID)

	switch existingToken.ID {
	case "":
		// no existing token -> insert
		query = `
			INSERT security_tokens
			SET
				id=?,
				user_id=?,
				token=?,
			    type=?,
				created_at=?,
				updated_at=?
		`

		_, err = r.DB.Exec(query,
			token.ID,
			token.UserID,
			token.Token,
			token.Type,
			token.CreatedAt,
			token.UpdatedAt,
		)
	default:
		// existing token -> update
		query = `
			UPDATE security_tokens
			SET
				token=?,
				updated_at=?
			WHERE id = ?
		`
		_, err = r.DB.Exec(query,
			token.Token,
			token.UpdatedAt,
			existingToken.ID,
		)
	}

	return err
}

// GetTokenByMetadata finds a auth.SecurityToken in the datastore
func (r *securityTokenRepository) GetTokenByMetadata(tokenMetadata *auth.TokenMetadata) (auth.SecurityToken, error) {
	var token auth.SecurityToken
	query := `
		SELECT 
			id, 
			user_id,
			token,
			type,
			created_at,
			updated_at
		FROM security_tokens 
		WHERE user_id = ? AND type = ? LIMIT 1
	`
	row := r.DB.QueryRow(query, tokenMetadata.UserID, tokenMetadata.Type)
	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.Type,
		&token.CreatedAt,
		&token.UpdatedAt)

	if err != nil {
		return auth.SecurityToken{}, terr.NewNotFoundError("token not found")
	}

	return token, nil
}

// RemoveTokenByMetadata removes a token from the datastore
func (r *securityTokenRepository) RemoveTokenByMetadata(tokenMetadata *auth.TokenMetadata) error {
	query := `DELETE FROM security_tokens WHERE user_id = ? AND type = ?`
	_, err := r.DB.Exec(query,
		tokenMetadata.UserID,
		tokenMetadata.Type,
	)
	return err
}
