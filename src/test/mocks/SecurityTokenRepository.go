// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	auth "sherman/src/domain/auth"

	mock "github.com/stretchr/testify/mock"
)

// SecurityTokenRepository is an autogenerated mock type for the SecurityTokenRepository type
type SecurityTokenRepository struct {
	mock.Mock
}

// CreateOrUpdateToken provides a mock function with given fields: token
func (_m *SecurityTokenRepository) CreateOrUpdateToken(token *auth.SecurityToken) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(*auth.SecurityToken) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTokenByMetadata provides a mock function with given fields: tokenMetadata
func (_m *SecurityTokenRepository) GetTokenByMetadata(tokenMetadata *auth.TokenMetadata) (auth.SecurityToken, error) {
	ret := _m.Called(tokenMetadata)

	var r0 auth.SecurityToken
	if rf, ok := ret.Get(0).(func(*auth.TokenMetadata) auth.SecurityToken); ok {
		r0 = rf(tokenMetadata)
	} else {
		r0 = ret.Get(0).(auth.SecurityToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*auth.TokenMetadata) error); ok {
		r1 = rf(tokenMetadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTokenByMetadata provides a mock function with given fields: tokenMetadata
func (_m *SecurityTokenRepository) RemoveTokenByMetadata(tokenMetadata *auth.TokenMetadata) error {
	ret := _m.Called(tokenMetadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(*auth.TokenMetadata) error); ok {
		r0 = rf(tokenMetadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}