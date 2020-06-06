package validator

import (
	"sherman/src/domain/auth"
)

type (
	// Validator validator.Validator interface definition
	Validator interface {
		ValidateUserParams(user *auth.User, action string) map[string]string
	}

	service struct{}
)

// New returns an instance of validator.Validator
func New() Validator {
	return &service{}
}
