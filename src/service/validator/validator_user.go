package validator

import (
	"sherman/src/domain/auth"
	"strings"
)

// ValidateUserParams validates /user/[route] route params, retrieves error messages for no compliant fields
func (s *service) ValidateUserParams(user *auth.User, action string) map[string]string {
	var errorMessages = make(map[string]string)

	const (
		firstNameRequired = "first_name is required"
		lastNameRequired  = "last_name is required"
		passwordRequired  = "password is required"
		emailRequired     = "email_address is required"
	)

	switch strings.ToLower(action) {
	case "register":
		if user.FirstName == "" {
			errorMessages["first_name_required"] = firstNameRequired
		}
		if user.LastName == "" {
			errorMessages["last_name_required"] = lastNameRequired
		}
		if user.Password == "" {
			errorMessages["password_required"] = passwordRequired
		}
		if user.EmailAddress == "" {
			errorMessages["email_address_required"] = emailRequired
		}
	case "login":
		if user.EmailAddress == "" {
			errorMessages["email_address_required"] = emailRequired
		}
		if user.Password == "" {
			errorMessages["password_required"] = passwordRequired
		}
	}
	return errorMessages
}
