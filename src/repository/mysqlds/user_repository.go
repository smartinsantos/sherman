package mysqlds

import (
	"database/sql"
	"sherman/src/domain/auth"
	"sherman/src/utils/exception"
	"strings"
)

// UserRepository sql implementation of auth.UserRepository
type UserRepository struct {
	DB *sql.DB
}

// CreateUser persist a auth.User from the datastore
func (r *UserRepository) CreateUser(user *auth.User) error {
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
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			err = exception.NewDuplicateEntryError("user already exist")
		}
		return err
	}

	return nil
}

// GetUserByID gets a auth.User by id in the datastore
func (r *UserRepository) GetUserByID(id string) (auth.User, error) {
	var user auth.User

	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	row := r.DB.QueryRow(query, id)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.EmailAddress,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "no rows") {
			err = exception.NewNotFoundError("user not found")
		}
		return auth.User{}, err
	}

	return user, nil
}

// GetUserByEmail gets a auth.User by email from the datastore
func (r *UserRepository) GetUserByEmail(email string) (auth.User, error) {
	var user auth.User

	query := `SELECT * FROM users WHERE email_address = ? LIMIT 1`
	row := r.DB.QueryRow(query, email)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.EmailAddress,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "no rows") {
			err = exception.NewNotFoundError("user not found")
		}
		return auth.User{}, err
	}

	return user, nil
}
