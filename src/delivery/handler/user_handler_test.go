package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"sherman/mocks"
	"sherman/src/domain/auth"
	// "sherman/src/domain/auth"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {


	t.Run("it should succeed", func(t *testing.T) {
		mockUserUseCase := new(mocks.UserUseCase)
		mocksSecurityTokenUseCase := new(mocks.SecurityTokenUseCase)
		mockValidatorService := new(mocks.Validator)
		mockValidatorService.
			On("ValidateUserParams", mock.Anything).
			Return(make(map[string]string))
		mockSecurityService := new(mocks.Security)
		mockPresenterService := new(mocks.Presenter)
		userHandler := NewUserHandler(
			mockUserUseCase,
			mocksSecurityTokenUseCase,
			mockValidatorService,
			mockSecurityService,
			mockPresenterService,
		)

		mockUser := auth.User{
			FirstName:    "first",
			LastName:     "last",
			EmailAddress: "some@email.com",
			Password:     "some_password",
		}

		userJson, err := json.Marshal(mockUser)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/api/v1/users/register", strings.NewReader(string(userJson)))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = userHandler.Register(ctx)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}