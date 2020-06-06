package presenter

import (
	"sherman/src/domain/auth"
	"sync"
)

type (
	Service interface {
		PresentUser(user *auth.User) auth.PresentedUser
	}

	service struct{}
)

var (
	instance Service
	once     sync.Once
)

// Get returns an instance of presenter.Service
func Get() Service {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}

// PresentUser returns a map of public auth.User keys, values
func (s *service) PresentUser(user *auth.User) auth.PresentedUser {
	return auth.PresentedUser{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Active:       user.Active,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
