package validator

import (
	"github.com/stretchr/testify/assert"
	"sherman/src/domain/auth"
	"testing"
)

func TestValidateUserParams(t *testing.T) {
	vs := New()
	mockUser := auth.User{
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "has a password",
	}

	errors := vs.ValidateUserParams(&mockUser, "register")
	assert.Equal(t, map[string]string{}, errors)
	errors = vs.ValidateUserParams(&mockUser, "login")
	assert.Equal(t, map[string]string{}, errors)

	mockUser = auth.User{
		FirstName:    "",
		LastName:     "",
		EmailAddress: "",
		Password:     "",
	}
	errors = vs.ValidateUserParams(&mockUser, "register")
	expected := map[string]string{
		"email_address_required": "email_address is required",
		"first_name_required":    "first_name is required",
		"last_name_required":     "last_name is required",
		"password_required":      "password is required",
	}
	assert.Equal(t, expected, errors)
	errors = vs.ValidateUserParams(&mockUser, "login")
	expected = map[string]string{
		"email_address_required": "email_address is required",
		"password_required":      "password is required",
	}
	assert.Equal(t, expected, errors)
}
