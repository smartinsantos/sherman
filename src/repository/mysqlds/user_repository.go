package mysqlds

import (
	"database/sql"
	"root/src/domain"
)

// UserRepository sql implementation of domain.UserRepository
type UserRepository struct {
	DB *sql.DB
}

// CreateUser persist a domain.user in the db
func (r *UserRepository) CreateUser(user *domain.User) error {
	query := `
		INSERT users
		SET
			id=?,
		    first_name=?,
			last_name=?,
			email_address=?,
			password=?,
			active=?,
			created_at=?,
			updated_at=?
	`

	_, err := r.DB.Exec(query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.EmailAddress,
		user.Password,
		user.Active,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail find a domain.user by email in the db
func (r *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	var err error
	var user domain.User

	query := `SELECT * FROM users WHERE email_address = ? LIMIT 1`
	row := r.DB.QueryRow(query, email)
	err = row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.EmailAddress,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}