package presenter

import (
	"sherman/src/domain/auth"
)

type (
	// Presenter presenter.Presenter interface definition
	Presenter interface {
		PresentUser(user *auth.User) auth.PresentedUser
	}

	service struct{}
)

// New returns an instance of presenter.Presenter
func New() Presenter {
	return &service{}
}
