package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sherman/mocks"
	"sherman/src/app/utils/exception"
	"sherman/src/domain/auth"
	"testing"
)

type userUseCaseMockDeps struct {
	userRepository  *mocks.UserRepository
	securityService *mocks.Security
}

func genUserUseCase() (auth.UserUseCase, userUseCaseMockDeps) {
	uucDeps := userUseCaseMockDeps{
		userRepository:  new(mocks.UserRepository),
		securityService: new(mocks.Security),
	}

	uuc := NewUserUseCase(
		uucDeps.userRepository,
		uucDeps.securityService,
	)

	return uuc, uucDeps
}

func TestRegister(t *testing.T) {
	mockUser := auth.User{
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
	}
	mockHashPassword := []byte("some-hashed-password")

	t.Run("it should succeed", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		muCopy := mockUser
		uucDeps.userRepository.On("CreateUser", mock.Anything).Return(nil)
		uucDeps.securityService.
			On("Hash", mock.AnythingOfType("string")).
			Return(mockHashPassword, nil)

		err := uuc.Register(&muCopy)

		assert.NoError(t, err)
		assert.NotEmpty(t, muCopy.ID)
		assert.NotEmpty(t, muCopy.CreatedAt)
		assert.NotEmpty(t, muCopy.UpdatedAt)
		assert.EqualValues(t, string(mockHashPassword), muCopy.Password)
		assert.EqualValues(t, true, muCopy.Active)
		assert.EqualValues(t, mockUser.FirstName, muCopy.FirstName)
		assert.EqualValues(t, mockUser.LastName, muCopy.LastName)
		assert.EqualValues(t, mockUser.EmailAddress, muCopy.EmailAddress)
	})

	t.Run("it should return an error", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		muCopy := mockUser
		mockError := errors.New("test register error")
		uucDeps.userRepository.On("CreateUser", mock.Anything).Return(nil)
		uucDeps.securityService.On("Hash", mock.AnythingOfType("string")).Return(nil, mockError)

		err := uuc.Register(&muCopy)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		muCopy := mockUser
		mockError := errors.New("test register error")
		uucDeps.securityService.
			On("Hash", mock.AnythingOfType("string")).
			Return(mockHashPassword, nil)
		uucDeps.userRepository.
			On("CreateUser", mock.Anything).
			Return(mockError)

		err := uuc.Register(&muCopy)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}

func TestVerifyCredentials(t *testing.T) {
	mockPassword := "some-password"
	mockHashedPassword := "some-hashed-password"
	mockUserRecord := auth.User{
		Password: mockHashedPassword,
	}
	mockUser := auth.User{
		Password: mockPassword,
	}

	t.Run("it should succeed", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		uucDeps.userRepository.On("GetUserByEmail", mock.Anything).Return(mockUserRecord, nil)
		uucDeps.securityService.
			On("VerifyPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(nil)

		userRecord, err := uuc.VerifyCredentials(&mockUser)

		assert.NoError(t, err)
		assert.EqualValues(t, mockUserRecord, userRecord)
	})

	t.Run("it should return an error", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		mockError := errors.New("get user by email error")
		uucDeps.userRepository.On("GetUserByEmail", mock.Anything).Return(auth.User{}, mockError)

		_, err := uuc.VerifyCredentials(&mockUser)

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("it should return an un-authorized error", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		mockError := exception.NewUnAuthorizedError("password doesn't match")
		uucDeps.securityService.
			On("VerifyPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(errors.New("some-error"))
		uucDeps.userRepository.On("GetUserByEmail", mock.Anything).Return(mockUserRecord, nil)

		_, err := uuc.VerifyCredentials(&mockUser)

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	mockUser := auth.User{
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
	}

	t.Run("it should succeed", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		uucDeps.userRepository.
			On("GetUserByID", mock.AnythingOfType("string")).
			Return(mockUser, nil)

		userRecord, err := uuc.GetUserByID("some-id")

		assert.NoError(t, err)
		assert.EqualValues(t, mockUser, userRecord)
	})

	t.Run("it should return an error", func(t *testing.T) {
		uuc, uucDeps := genUserUseCase()
		mockError := errors.New("get user by id error")
		uucDeps.userRepository.
			On("GetUserByID", mock.Anything).
			Return(auth.User{}, mockError)

		_, err := uuc.GetUserByID("some-id")

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}
