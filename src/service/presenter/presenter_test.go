package presenter

import (
	"github.com/stretchr/testify/assert"
	_ "sherman/src/app/testing"
	"sherman/src/domain/auth"
	"testing"
	"time"
)

func TestPresentUser(t *testing.T) {
	ps := New()
	mockUser := auth.User{
		ID:           "some-id",
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "has a password",
		Active:       true,
		CreatedAt:    time.Unix(0, 0),
		UpdatedAt:    time.Unix(0, 0),
	}

	expected := auth.PresentedUser{
		ID:           mockUser.ID,
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		EmailAddress: mockUser.EmailAddress,
		Active:       mockUser.Active,
		CreatedAt:    mockUser.CreatedAt,
		UpdatedAt:    mockUser.UpdatedAt,
	}
	actual := ps.PresentUser(&mockUser)

	assert.Equal(t, expected, actual)
}
