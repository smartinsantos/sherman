// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	auth "sherman/src/domain/auth"

	mock "github.com/stretchr/testify/mock"
)

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// ValidateUserParams provides a mock function with given fields: user, action
func (_m *Validator) ValidateUserParams(user *auth.User, action string) map[string]string {
	ret := _m.Called(user, action)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(*auth.User, string) map[string]string); ok {
		r0 = rf(user, action)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}