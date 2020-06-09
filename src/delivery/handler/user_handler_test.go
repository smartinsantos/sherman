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
	"time"
)

type userHandlerMockDeps struct {
	userUseCase          *mocks.UserUseCase
	securityTokenUseCase *mocks.SecurityTokenUseCase
	validatorService     *mocks.Validator
	securityService      *mocks.Security
	presenterService     *mocks.Presenter
}

func genMockUserHandler() (UserHandler, userHandlerMockDeps) {
	uhDeps := userHandlerMockDeps{
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
		uh, uhDeps := genMockUserHandler()
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

		if assert.NoError(t, uh.Register(ctx)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "{\"data\":null}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, _ := genMockUserHandler()
		userJSON, err := json.Marshal("wrong-params")
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Register(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
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

		if assert.NoError(t, uh.Register(ctx)) {
			assert.EqualValues(t, http.StatusUnprocessableEntity, rec.Code)
			assert.Equal(t, "{\"data\":null,\"errors\":{\"some-error\":\"some error\"}}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
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

		if assert.NoError(t, uh.Register(ctx)) {
			assert.EqualValues(t, http.StatusForbidden, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"register error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
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

		if assert.NoError(t, uh.Register(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})
}

func TestLogin(t *testing.T) {
	mockUser := auth.User{
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some_password",
	}

	mockToken := auth.SecurityToken{
		ID:        "some-id",
		UserID:    "some-user-id",
		Token:     "some-token",
		Type:      "some-token-type",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t.Run("it should succeed", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		uhDeps.userUseCase.
			On("VerifyCredentials", mock.Anything).
			Return(mockUser, nil)
		uhDeps.securityTokenUseCase.
			On("GenAccessToken", mock.AnythingOfType("string")).
			Return(mockToken, nil)
		uhDeps.securityTokenUseCase.
			On("GenRefreshToken", mock.AnythingOfType("string")).
			Return(mockToken, nil)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/login", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "REFRESH_TOKEN=some-token; Path=/; Max-Age=3600; HttpOnly", rec.Header().Get("Set-Cookie"))
			assert.Equal(t, "{\"data\":{\"access_token\":\"some-token\"}}\n", rec.Body.String())
		}
	})
}
