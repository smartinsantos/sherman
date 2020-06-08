package handler

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"sherman/mocks"
	"sherman/src/app/utils/exception"
	"sherman/src/domain/auth"
	"strings"
	"testing"
)

type mockDependencies struct {
	userUseCase          *mocks.UserUseCase
	securityTokenUseCase *mocks.SecurityTokenUseCase
	validatorService     *mocks.Validator
	securityService      *mocks.Security
	presenterService     *mocks.Presenter
}

func generateUserHandler() (UserHandler, mockDependencies) {
	uhDeps := mockDependencies{
		userUseCase:          new(mocks.UserUseCase),
		securityTokenUseCase: new(mocks.SecurityTokenUseCase),
		validatorService:     new(mocks.Validator),
		securityService:      new(mocks.Security),
		presenterService:     new(mocks.Presenter),
	}

	uh := NewUserHandler(
		uhDeps.userUseCase,
		uhDeps.securityTokenUseCase,
		uhDeps.validatorService,
		uhDeps.securityService,
		uhDeps.presenterService,
	)

	return uh, uhDeps
}

func TestRegister(t *testing.T) {
	mockUser := auth.User{
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some_password",
	}

	t.Run("it should succeed", func(t *testing.T) {
		uh, uhDeps := generateUserHandler()
		uhDeps.userUseCase.On("Register", mock.Anything).Return(nil)
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = uh.Register(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, http.StatusCreated, rec.Code)
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, _ := generateUserHandler()
		userJSON, err := json.Marshal("wrong-params")
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = uh.Register(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := generateUserHandler()
		mockErrorMessages := make(map[string]string)
		mockErrorMessages["some-error"] = "some error"
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(mockErrorMessages)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = uh.Register(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := generateUserHandler()
		mockError := exception.NewDuplicateEntryError("register error")
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		uhDeps.userUseCase.On("Register", mock.Anything).Return(mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = uh.Register(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, http.StatusForbidden, rec.Code)
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := generateUserHandler()
		mockError := errors.New("some-error")
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		uhDeps.userUseCase.On("Register", mock.Anything).Return(mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = uh.Register(ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	})
}
