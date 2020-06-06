package validator

import (
	"sherman/src/domain/auth"
	"sync"
)

type (
	// Service validator.ValidatorService interface definition
	ValidatorService interface {
		ValidateUserParams(user *auth.User, action string) map[string]string
	}

	service struct{}
)

var (
	instance ValidatorService
	once     sync.Once
)

// Get returns an instance of validator.ValidatorService
func Get() ValidatorService {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}
