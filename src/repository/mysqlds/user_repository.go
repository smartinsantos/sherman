package mysqlds

import (
	"database/sql"

	"root/src/domain"
)

type dsUserRepository struct {
	db *sql.DB
}

// NewDsUserRepository creates a new object representation of domain.UserRepository interface
func NewDsUserRepository(db *sql.DB) (domain.UserRepository, error) {
	repository := dsUserRepository { db: db }

	return &repository, nil
}

// Creates a user
func (r *dsUserRepository) CreateUser(user *domain.User) error {
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

	_, err := r.db.Exec(query,
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

// Find user by email
func (r *dsUserRepository) GetUserByEmail(email string) (domain.User, error) {
	var err error
	var user domain.User

	query := `SELECT * FROM users WHERE email_address = ? LIMIT 1`
	row := r.db.QueryRow(query, email)
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