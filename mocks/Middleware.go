// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// JWT provides a mock function with given fields:
func (_m *Middleware) JWT() echo.MiddlewareFunc {
	ret := _m.Called()

	var r0 echo.MiddlewareFunc
	if rf, ok := ret.Get(0).(func() echo.MiddlewareFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.MiddlewareFunc)
		}
	}

	return r0
}

// ZeroLog provides a mock function with given fields:
func (_m *Middleware) ZeroLog() echo.MiddlewareFunc {
	ret := _m.Called()

	var r0 echo.MiddlewareFunc
	if rf, ok := ret.Get(0).(func() echo.MiddlewareFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.MiddlewareFunc)
		}
	}

	return r0
}
