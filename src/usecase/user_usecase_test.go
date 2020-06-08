package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sherman/mocks"
	"sherman/src/app/utils/exception"
	"sherman/src/domain/auth"
	"sherman/src/service/security"
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

	t.Run("it should succeed", func(t *testing.T) {
		// uuc, uucDeps := genUserUseCase()

		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("CreateUser", mock.Anything).Return(nil)
		muCopy := mockUser
		securityService := security.New()
		userUseCase := NewUserUseCase(mockUserRepo, securityService)

		err := userUseCase.Register(&muCopy)

		assert.NoError(t, err)
		assert.NotEmpty(t, muCopy.ID)
		assert.NotEmpty(t, muCopy.CreatedAt)
		assert.NotEmpty(t, muCopy.UpdatedAt)
		assert.EqualValues(t, true, muCopy.Active)
		assert.EqualValues(t, mockUser.FirstName, muCopy.FirstName)
		assert.EqualValues(t, mockUser.LastName, muCopy.LastName)
		assert.EqualValues(t, mockUser.EmailAddress, muCopy.EmailAddress)
		err = securityService.VerifyPassword(muCopy.Password, mockUser.Password)
		assert.NoError(t, err)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		muCopy := mockUser
		mockError := errors.New("test register error")
		mockSecurityService := new(mocks.Security)
		userUseCase := NewUserUseCase(mockUserRepo, mockSecurityService)

		mockUserRepo.On("CreateUser", mock.Anything).Return(nil)
		mockSecurityService.On("Hash", mock.AnythingOfType("string")).Return(nil, mockError)

		err := userUseCase.Register(&muCopy)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		muCopy := mockUser
		mockError := errors.New("test register error")
		securityService := security.New()
		userUseCase := NewUserUseCase(mockUserRepo, securityService)

		mockUserRepo.On("CreateUser", mock.Anything).Return(mockError)

		err := userUseCase.Register(&muCopy)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}

func TestVerifyCredentials(t *testing.T) {
	securityService := security.New()
	mockPassword := "some-password"
	hashPassword, err := securityService.Hash(mockPassword)
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

		userUseCase := NewUserUseCase(mockUserRepo, securityService)
		userRecord, err := userUseCase.VerifyCredentials(&mockUser)

		assert.NoError(t, err)
		assert.EqualValues(t, mockUserRecord, userRecord)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockError := errors.New("get user by email error")
		mockUserRepo.On("GetUserByEmail", mock.Anything).Return(auth.User{}, mockError)
		securityService := security.New()
		userUseCase := NewUserUseCase(mockUserRepo, securityService)
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
		securityService := security.New()
		userUseCase := NewUserUseCase(mockUserRepo, securityService)
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
		securityService := security.New()
		userUseCase := NewUserUseCase(mockUserRepo, securityService)
		userRecord, err := userUseCase.GetUserByID("some-id")

		assert.NoError(t, err)
		assert.EqualValues(t, mockUser, userRecord)
	})

	t.Run("it should return an error", func(t *testing.T) {
		mockError := errors.New("get user by id error")
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("GetUserByID", mock.Anything).Return(auth.User{}, mockError)
		securityService := security.New()
		userUseCase := NewUserUseCase(mockUserRepo, securityService)
		_, err := userUseCase.GetUserByID("some-id")

		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})
}
