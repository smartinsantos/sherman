package presenter

import "sherman/src/domain/auth"

// PresentUser returns a map of public auth.User keys, values
func (s *service) PresentUser(user *auth.User) auth.PresentedUser {
	return auth.PresentedUser{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Active:       user.Active,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
