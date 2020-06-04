package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sherman/src/domain/auth"
	"sherman/src/test/mocks"
	"sherman/src/usecase"
	"sherman/src/utils/security"
	"testing"
)

func TestRegister(t *testing.T) {
	mockUser := auth.User{
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
	}
	mockUserRepo := new(mocks.UserRepository)
	var userUseCase auth.UserUseCase = &usecase.UserUseCase{
		UserRepo: mockUserRepo,
	}

	t.Run("success", func(t *testing.T) {
		muCopy := mockUser
		mockUserRepo.On("Register", mock.Anything).Return(nil)
		mockUserRepo.On("CreateUser", mock.Anything).Return(nil)

		err := userUseCase.Register(&muCopy)

		assert.NoError(t, err)
		assert.NotEmpty(t, muCopy.ID)
		assert.NotEmpty(t, muCopy.CreatedAt)
		assert.NotEmpty(t, muCopy.UpdatedAt)
		assert.EqualValues(t, true, muCopy.Active)
		assert.EqualValues(t, mockUser.FirstName, muCopy.FirstName)
		assert.EqualValues(t, mockUser.LastName, muCopy.LastName)
		assert.EqualValues(t, mockUser.EmailAddress, muCopy.EmailAddress)
		err = security.VerifyPassword(muCopy.Password, mockUser.Password)
		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		muCopy := mockUser
		mockError := errors.New("TestRegister Error")
		mockUserRepo.On("Register", mock.Anything).Return(mockError)
		mockUserRepo.On("CreateUser", mock.Anything).Return(mockError)

		err := userUseCase.Register(&muCopy)
		assert.Error(t, mockError, err)
	})
}

func TestVerifyCredentials(t *testing.T) {}

func TestGetUserByID(t *testing.T) {}
