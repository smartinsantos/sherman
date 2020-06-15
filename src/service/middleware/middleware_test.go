package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"sherman/mocks"
	"sherman/src/domain/auth"
	"testing"
)

type middlewareMockDeps struct {
	securityService *mocks.Security
}

func genMockMiddleware() (Middleware, middlewareMockDeps) {
	mDeps := middlewareMockDeps{
		securityService: new(mocks.Security),
	}
	m := New(mDeps.securityService)
	return m, mDeps
}

func TestUserAuthMiddleware(t *testing.T) {
	t.Run("request should go thru", func(t *testing.T) {
		m, mDeps := genMockMiddleware()
		mDeps.securityService.
			On("GetAndValidateAccessToken", mock.Anything).
			Return(auth.TokenMetadata{}, nil)

		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}

		h := m.UserAuthMiddleware()(handler)

		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "test", res.Body.String())
		}
	})

	t.Run("request should not go thru", func(t *testing.T) {
		m, mDeps := genMockMiddleware()
		mDeps.securityService.
			On("GetAndValidateAccessToken", mock.Anything).
			Return(auth.TokenMetadata{}, errors.New("some error"))

		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}

		h := m.UserAuthMiddleware()(handler)

		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"invalid token\"}\n", res.Body.String())
		}
	})
}
