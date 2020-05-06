package validator

import (
	"github.com/smartinsantos/go-auth-api/domain"
	"strings"
)

func ValidateUserParams(user *domain.User, action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
		default:
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
	}
	return errorMessages
}