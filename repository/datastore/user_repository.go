package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/smartinsantos/go-auth-api/domain"
)

type dsUserRepository struct {
	db *gorm.DB
}

// NewDsUserRepository creates a new object representation of domain.UserRepository interface
func NewDsUserRepository(db *gorm.DB) domain.UserRepository {
	return &dsUserRepository{ db: db }
}

// Creates a user
func (ur *dsUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}