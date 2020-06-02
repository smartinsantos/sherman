// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	auth "sherman/src/domain/auth"

	mock "github.com/stretchr/testify/mock"
)

// SecurityTokenUseCase is an autogenerated mock type for the SecurityTokenUseCase type
type SecurityTokenUseCase struct {
	mock.Mock
}

// GenAccessToken provides a mock function with given fields: userID
func (_m *SecurityTokenUseCase) GenAccessToken(userID string) (auth.SecurityToken, error) {
	ret := _m.Called(userID)

	var r0 auth.SecurityToken
	if rf, ok := ret.Get(0).(func(string) auth.SecurityToken); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(auth.SecurityToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenRefreshToken provides a mock function with given fields: userID
func (_m *SecurityTokenUseCase) GenRefreshToken(userID string) (auth.SecurityToken, error) {
	ret := _m.Called(userID)

	var r0 auth.SecurityToken
	if rf, ok := ret.Get(0).(func(string) auth.SecurityToken); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(auth.SecurityToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRefreshTokenStored provides a mock function with given fields: refreshTokenMetadata
func (_m *SecurityTokenUseCase) IsRefreshTokenStored(refreshTokenMetadata *auth.TokenMetadata) bool {
	ret := _m.Called(refreshTokenMetadata)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*auth.TokenMetadata) bool); ok {
		r0 = rf(refreshTokenMetadata)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveRefreshToken provides a mock function with given fields: refreshTokenMetadata
func (_m *SecurityTokenUseCase) RemoveRefreshToken(refreshTokenMetadata *auth.TokenMetadata) error {
	ret := _m.Called(refreshTokenMetadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(*auth.TokenMetadata) error); ok {
		r0 = rf(refreshTokenMetadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
