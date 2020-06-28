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
	_ "sherman/src/app/testing"
	"sherman/src/app/utils/terr"
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
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
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
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
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
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
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
		mockError := terr.NewDuplicateEntryError("register error")
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		uhDeps.userUseCase.On("Register", mock.Anything).Return(mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
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
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
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
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
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

	t.Run("it should return error", func(t *testing.T) {
		uh, _ := genMockUserHandler()
		userJSON, err := json.Marshal("wrong-params")
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
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
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.EqualValues(t, http.StatusUnprocessableEntity, rec.Code)
			assert.Equal(t, "{\"data\":null,\"errors\":{\"some-error\":\"some error\"}}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		mockError := terr.NewNotFoundError("verify credentials not found error")
		uhDeps.userUseCase.
			On("VerifyCredentials", mock.Anything).
			Return(auth.User{}, mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"verify credentials not found error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		mockError := terr.NewUnAuthorizedError("verify credentials unauthorized error")
		uhDeps.userUseCase.
			On("VerifyCredentials", mock.Anything).
			Return(auth.User{}, mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"verify credentials unauthorized error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		mockError := errors.New("any verify credentials error")
		uhDeps.userUseCase.
			On("VerifyCredentials", mock.Anything).
			Return(auth.User{}, mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.validatorService.
			On("ValidateUserParams", mock.Anything, mock.AnythingOfType("string")).
			Return(make(map[string]string))
		uhDeps.userUseCase.
			On("VerifyCredentials", mock.Anything).
			Return(mockUser, nil)
		mockError := errors.New("generate access token error")
		uhDeps.securityTokenUseCase.
			On("GenAccessToken", mock.AnythingOfType("string")).
			Return(auth.SecurityToken{}, mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
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
		mockError := errors.New("generate refresh token error")
		uhDeps.securityTokenUseCase.
			On("GenRefreshToken", mock.AnythingOfType("string")).
			Return(auth.SecurityToken{}, mockError)

		userJSON, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(string(userJSON)))
		assert.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Login(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})
}

func TestRefreshAccessToken(t *testing.T) {
	mockTokenMeta := auth.TokenMetadata{
		UserID: "some-user-id",
		Type:   "some-token-type",
		Token:  "some-token",
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
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(mockTokenMeta, nil)
		uhDeps.securityTokenUseCase.
			On("IsRefreshTokenStored", mock.Anything).
			Return(true)
		uhDeps.securityTokenUseCase.
			On("GenAccessToken", mock.AnythingOfType("string")).
			Return(mockToken, nil)

		e := echo.New()
		req, err := http.NewRequest(echo.PATCH, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.RefreshAccessToken(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "{\"data\":{\"access_token\":\"some-token\"}}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		mockError := errors.New("get and validate refresh token error")
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(auth.TokenMetadata{}, mockError)

		e := echo.New()
		req, err := http.NewRequest(echo.PATCH, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.RefreshAccessToken(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"invalid refresh token\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(mockTokenMeta, nil)
		uhDeps.securityTokenUseCase.
			On("IsRefreshTokenStored", mock.Anything).
			Return(false)

		e := echo.New()
		req, err := http.NewRequest(echo.PATCH, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.RefreshAccessToken(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"invalid refresh token\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(mockTokenMeta, nil)
		uhDeps.securityTokenUseCase.
			On("IsRefreshTokenStored", mock.Anything).
			Return(true)
		mockError := errors.New("gen access token error")
		uhDeps.securityTokenUseCase.
			On("GenAccessToken", mock.AnythingOfType("string")).
			Return(auth.SecurityToken{}, mockError)

		e := echo.New()
		req, err := http.NewRequest(echo.PATCH, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.RefreshAccessToken(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"gen access token error\"}\n", rec.Body.String())
		}
	})
}

func TestGetUser(t *testing.T) {
	lo, _ := time.LoadLocation("UTC")
	mockUser := auth.User{
		ID:           "some-id",
		FirstName:    "first",
		LastName:     "last",
		EmailAddress: "some@email.com",
		Password:     "some-password",
		Active:       true,
		CreatedAt:    time.Unix(0, 0).In(lo),
		UpdatedAt:    time.Unix(0, 0).In(lo),
	}

	t.Run("it should succeed", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.userUseCase.
			On("GetUserByID", mock.Anything).
			Return(mockUser, nil)
		uhDeps.presenterService.
			On("PresentUser", mock.Anything).
			Return(auth.PresentedUser{
				ID:           mockUser.ID,
				FirstName:    mockUser.FirstName,
				LastName:     mockUser.LastName,
				EmailAddress: mockUser.EmailAddress,
				Active:       mockUser.Active,
				CreatedAt:    mockUser.CreatedAt,
				UpdatedAt:    mockUser.UpdatedAt,
			})

		e := echo.New()
		req, err := http.NewRequest(
			echo.GET, "/some-url/"+mockUser.ID,
			strings.NewReader(""),
		)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("some-url/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues(mockUser.ID)

		if assert.NoError(t, uh.GetUser(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(
				t,
				"{\"data\":{\"user\":{\"id\":\"some-id\",\"first_name\":\"first\",\"last_name\":\"last\",\"email_address\":\"some@email.com\",\"active\":true,\"created_at\":\"1970-01-01T00:00:00Z\",\"updated_at\":\"1970-01-01T00:00:00Z\"}}}\n",
				rec.Body.String(),
			)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		mockError := terr.NewNotFoundError("get user by id not found error")
		uhDeps.userUseCase.
			On("GetUserByID", mock.Anything).
			Return(auth.User{}, mockError)

		e := echo.New()
		req, err := http.NewRequest(
			echo.GET, "/some-url/"+mockUser.ID,
			strings.NewReader(""),
		)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("some-url/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues(mockUser.ID)

		if assert.NoError(t, uh.GetUser(ctx)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"get user by id not found error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		mockError := errors.New("any get user by id error")
		uhDeps.userUseCase.
			On("GetUserByID", mock.Anything).
			Return(auth.User{}, mockError)

		e := echo.New()
		req, err := http.NewRequest(
			echo.GET, "/some-url/"+mockUser.ID,
			strings.NewReader(""),
		)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("some-url/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues(mockUser.ID)

		if assert.NoError(t, uh.GetUser(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})
}

func TestLogout(t *testing.T) {
	mockTokenMeta := auth.TokenMetadata{
		UserID: "some-user-id",
		Type:   "some-token-type",
		Token:  "some-token",
	}

	t.Run("it should succeed", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(mockTokenMeta, nil)
		uhDeps.securityTokenUseCase.
			On("RemoveRefreshToken", mock.Anything).
			Return(nil)

		e := echo.New()
		req, err := http.NewRequest(echo.DELETE, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Logout(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "REFRESH_TOKEN=; Path=/; HttpOnly", rec.Header().Get("Set-Cookie"))
			assert.Equal(t, "{\"data\":null}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		mockError := errors.New("get and validate refresh token error")
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(auth.TokenMetadata{}, mockError)

		e := echo.New()
		req, err := http.NewRequest(echo.PATCH, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Logout(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"get and validate refresh token error\"}\n", rec.Body.String())
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		uh, uhDeps := genMockUserHandler()
		uhDeps.securityService.
			On("GetAndValidateRefreshToken", mock.Anything).
			Return(mockTokenMeta, nil)
		uhDeps.securityTokenUseCase.
			On("RemoveRefreshToken", mock.Anything).
			Return(errors.New("remove refresh token error"))

		e := echo.New()
		req, err := http.NewRequest(echo.PATCH, "/some-url", strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, uh.Logout(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"internal server error\"}\n", rec.Body.String())
		}
	})
}
