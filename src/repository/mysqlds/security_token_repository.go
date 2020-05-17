package mysqlds

import (
	"database/sql"
	"root/src/domain/auth"
)

// SecurityTokenRepository sql implementation of auth.SecurityTokenRepository
type SecurityTokenRepository struct {
	DB *sql.DB
}

// CreateOrUpdateToken persist a auth.SecurityToken in the db
func (r *SecurityTokenRepository) CreateOrUpdateToken(token *auth.SecurityToken) error {
	var err error
	var query string
	var exitingTokenId string

	// find token id if it exist
	query = `SELECT id FROM security_tokens WHERE user_id = ? LIMIT 1`
	row := r.DB.QueryRow(query, token.UserID)
	_ = row.Scan(&exitingTokenId)

	// no existing token -> insert
	if len(exitingTokenId) == 0 {
		query = `
			INSERT security_tokens
			SET
				id=?,
				user_id=?,
				token=?,
				created_at=?,
				updated_at=?
		`

		_, err = r.DB.Exec(query,
			token.ID,
			token.UserID,
			token.Token,
			token.CreatedAt,
			token.UpdatedAt,
		)

		return err
	}

	// existing token -> update
	query = `
		UPDATE security_tokens
		SET
			user_id=?,
			token=?,
			created_at=?,
			updated_at=?
		WHERE id = ?
	`
	_, err = r.DB.Exec(query,
		token.UserID,
		token.Token,
		token.CreatedAt,
		token.UpdatedAt,
		token.ID,
	)
	return err
}