package presenter

import (
	"sherman/src/domain/auth"
	"sync"
)

type (
	// PresenterService presenter.PresenterService interface definition
	PresenterService interface {
		PresentUser(user *auth.User) auth.PresentedUser
	}

	service struct{}
)

var (
	instance PresenterService
	once     sync.Once
)

// Get returns an instance of presenter.PresenterService
func Get() PresenterService {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}
