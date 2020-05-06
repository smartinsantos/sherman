package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/smartinsantos/go-auth-api/domain"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlArticleRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{ db: db }
}

// Creates a user
func (ur *mysqlUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}