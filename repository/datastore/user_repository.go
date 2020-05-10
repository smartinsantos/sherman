package datastore

import (
	"database/sql"
	"fmt"

	"github.com/gchaincl/dotsql"

	"root/domain"
)

type dsUserRepository struct {
	db *sql.DB
	dot *dotsql.DotSql
}

// NewDsUserRepository creates a new object representation of domain.UserRepository interface
func NewDsUserRepository(db *sql.DB) (domain.UserRepository, error) {
	dot, err := sqlLoader("user_repository")
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
	res, err := ur.dot.Exec(ur.db, "create-user")
	if err != nil {
		return nil, err
	}

	fmt.Println("res =>", res)

	return user, nil
}