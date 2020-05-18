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