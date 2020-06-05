package mysqlds

import (
	"database/sql"
	"sherman/src/app/utils/exception"
	"sherman/src/domain/auth"
	"strings"
)

// userRepository sql implementation of auth.UserRepository
type userRepository struct {
	DB *sql.DB
}

// NewUserRepository constructor
func NewUserRepository(db *sql.DB) auth.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) scanUserRow(row *sql.Row) (auth.User, error) {
	var user auth.User

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

// CreateUser persist a auth.User from the datastore
func (r *userRepository) CreateUser(user *auth.User) error {
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
func (r *userRepository) GetUserByID(id string) (auth.User, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			email_address,
			password,
			active,
			created_at,
			updated_at
		FROM users 
		WHERE id = ? LIMIT 1
	`
	row := r.DB.QueryRow(query, id)
	return r.scanUserRow(row)
}

// GetUserByEmail gets a auth.User by email from the datastore
func (r *userRepository) GetUserByEmail(email string) (auth.User, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			email_address,
			password,
			active,
			created_at,
			updated_at
		FROM users
		WHERE email_address = ? LIMIT 1
	`
	row := r.DB.QueryRow(query, email)
	return r.scanUserRow(row)
}
