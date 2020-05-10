package mysqlds

import (
	"database/sql"
	"github.com/gchaincl/dotsql"

	"root/src/domain"
)

type dsUserRepository struct {
	db *sql.DB
	dot *dotsql.DotSql
}

// NewDsUserRepository creates a new object representation of domain.UserRepository interface
func NewDsUserRepository(db *sql.DB) (domain.UserRepository, error) {
	dot, err := loadSqlDot("user_repository.sql")
	if err != nil {
		return nil, err
	}

	repository := dsUserRepository {
		db: db,
		dot: dot,
	}

	return &repository, nil
}

// Creates a user
func (ur *dsUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	_, err := ur.dot.Exec(
		ur.db,
		"create-user",
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
		return nil, err
	}

	return user, nil
}