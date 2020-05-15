package validator

import (
	"root/src/domain/auth"
	"strings"
)

// ValidateUserParams validates /user/[route] route params, retrieves error messages for no compliant fields
func ValidateUserParams(user *auth.User, action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
		case "register":
			if user.FirstName == "" {
				errorMessages["first_name_required"] = "First name is required"
			}
			if user.LastName == "" {
				errorMessages["last_name_required"] = "Last name is required"
			}
			if user.Password == "" {
				errorMessages["password_required"] = "Password is required"
			}
			if user.EmailAddress == "" {
				errorMessages["email_address_required"] = "Email Address is required"
			}
		case "login":
			if user.EmailAddress == "" {
				errorMessages["email_address_required"] = "Email Address is required"
			}
			if user.Password == "" {
				errorMessages["password_required"] = "Password is required"
			}
		default: // do nothing
	}
	return errorMessages
}