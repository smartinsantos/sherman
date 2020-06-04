package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sherman/src/domain/auth"
	"sherman/src/test/mocks"
	"sherman/src/utils/exception"
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

	t.Run("it should succeed", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("CreateUser", mock.Anything).Return(nil)
		muCopy := mockUser

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}

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

	t.Run("it should throw an error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		muCopy := mockUser
		mockError := errors.New("TestRegister Error")

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}

		mockUserRepo.On("CreateUser", mock.Anything).Return(mockError)

		err := userUseCase.Register(&muCopy)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}

func TestVerifyCredentials(t *testing.T) {
	mockPassword := "some-password"
	hashPassword, err := security.Hash(mockPassword)
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	mockHashedPassword := string(hashPassword)
	mockUserRecord := auth.User{
		Password: mockHashedPassword,
	}

	mockUser := auth.User{
		Password: mockPassword,
	}

	t.Run("it should succeed", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetUserByEmail", mock.Anything).Return(mockUserRecord, nil)

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}
		userRecord, err := userUseCase.VerifyCredentials(&mockUser)

		assert.NoError(t, err)
		assert.EqualValues(t, mockUserRecord, userRecord)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockError := errors.New("GetUserByEmail Error")
		mockUserRepo.On("GetUserByEmail", mock.Anything).Return(auth.User{}, mockError)

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}
		_, err := userUseCase.VerifyCredentials(&mockUser)

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("it should return an un-authorized error", func(t *testing.T) {
		mockError := exception.NewUnAuthorizedError("password doesn't match")
		mockUserWithWrongPassword := auth.User{
			Password: "wrong-password",
		}
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetUserByEmail", mock.Anything).Return(mockUserRecord, nil)

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}
		_, err := userUseCase.VerifyCredentials(&mockUserWithWrongPassword)

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("it should succeed", func(t *testing.T) {
		mockUser := auth.User{
			FirstName:    "first",
			LastName:     "last",
			EmailAddress: "some@email.com",
			Password:     "some-password",
		}

		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetUserByID", mock.AnythingOfType("string")).Return(mockUser, nil)

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}
		userRecord, err := userUseCase.GetUserByID("some-id")

		assert.NoError(t, err)
		assert.EqualValues(t, mockUser, userRecord)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockError := errors.New("GetUserByID error")
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetUserByID", mock.Anything).Return(auth.User{}, mockError)

		var userUseCase auth.UserUseCase = &UserUseCase{
			UserRepo: mockUserRepo,
		}
		_, err := userUseCase.GetUserByID("some-id")

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}
