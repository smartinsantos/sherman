package mysqlds

import (
	"database/sql"
	"root/src/domain/auth"
	"root/src/utils/exception"
	"strings"
)

// SecurityTokenRepository sql implementation of auth.SecurityTokenRepository
type SecurityTokenRepository struct {
	DB *sql.DB
}

// CreateToken persist a auth.SecurityToken in the db
func (r *SecurityTokenRepository) CreateToken(token *auth.SecurityToken) error {
	query := `
		INSERT security_tokens
		SET
			id=?,
		    user_id=?,
			token=?,
			created_at=?,
			updated_at=?
	`

	_, err := r.DB.Exec(query,
		token.ID,
		token.UserID,
		token.Token,
		token.CreatedAt,
		token.UpdatedAt,
	)

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			err = exception.NewDuplicateEntryError("token already exist")
		}
		return err
	}

	return nil
}