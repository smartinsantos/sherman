package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"sherman/src/app/config"
	_ "sherman/src/app/testing"
	"sherman/src/domain/auth"
	"strings"
	"testing"
	"time"
)

func TestValidateHash(t *testing.T) {
	mockPassword := "some-password"
	actualHash, err := New(config.Get()).Hash(mockPassword)
	if assert.NoError(t, err) {
		err = bcrypt.CompareHashAndPassword(actualHash, []byte(mockPassword))
		assert.NoError(t, err)
		err = bcrypt.CompareHashAndPassword(actualHash, []byte("some-other-password"))
		assert.Error(t, err)
	}
}

func TestVerifyPassword(t *testing.T) {
	mockPassword := "some-password"
	hash, err := bcrypt.GenerateFromPassword([]byte(mockPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	ss := New(config.Get())
	err = ss.VerifyPassword(string(hash), mockPassword)
	assert.NoError(t, err)
	err = ss.VerifyPassword(string(hash), "some-other-password")
	assert.Error(t, err)
}

func TestGenToken(t *testing.T) {
	mockUserID := "some-user-id"
	mockTokenType := "some-token-type"
	mockIat := time.Now().Unix()
	mockExp := time.Now().Add(time.Minute * time.Duration(15)).Unix()

	tokenStr, err := New(config.Get()).GenToken(mockUserID, mockTokenType, mockIat, mockExp)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, tokenStr)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid refresh token")
		}

		return []byte(config.Get().Jwt.Secret), nil
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	claims := token.Claims.(jwt.MapClaims)

	assert.EqualValues(t, claims["user_id"], mockUserID)
	assert.EqualValues(t, claims["type"], mockTokenType)
	assert.EqualValues(t, claims["iat"], mockIat)
	assert.EqualValues(t, claims["exp"], mockExp)
}

func TestGetAndValidateAccessToken(t *testing.T) {
	ss := New(config.Get())
	mockUserID := "some-user-id"
	mockTokenType := "some-token-type"
	mockIat := time.Now().Unix()
	mockExp := time.Now().Add(time.Minute * time.Duration(15)).Unix()

	t.Run("it should succeed", func(t *testing.T) {
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		mockTokenStr, err := New(config.Get()).GenToken(mockUserID, mockTokenType, mockIat, mockExp)
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		mockTokenMeta := auth.TokenMetadata{
			UserID: mockUserID,
			Type:   mockTokenType,
			Token:  mockTokenStr,
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+mockTokenStr)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateAccessToken(ctx)
		if assert.NoError(t, err) {
			assert.Equal(t, mockTokenMeta, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateAccessToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "access token not found", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// wrong signing method
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"user_id": mockUserID,
			"type":    mockTokenType,
			"iat":     mockIat,
			"exp":     mockExp,
		})
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		tokenStr, err := token.SignedString(key)
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+tokenStr)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateAccessToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// wrong secret
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": mockUserID,
			"type":    mockTokenType,
			"iat":     mockIat,
			"exp":     mockExp,
		})
		tokenStr, err := token.SignedString([]byte("some-other-secret"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+tokenStr)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateAccessToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// expired token
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": mockUserID,
			"type":    mockTokenType,
			"iat":     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			"exp":     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		})
		tokenStr, err := token.SignedString([]byte(config.Get().Jwt.Secret))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+tokenStr)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateAccessToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// no claims
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
		tokenStr, err := token.SignedString([]byte(config.Get().Jwt.Secret))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+tokenStr)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateAccessToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})
}

func TestGetAndValidateRefreshToken(t *testing.T) {
	ss := New(config.Get())
	mockUserID := "some-user-id"
	mockTokenType := "some-token-type"
	mockIat := time.Now().Unix()
	mockExp := time.Now().Add(time.Minute * time.Duration(15)).Unix()

	t.Run("it should succeed", func(t *testing.T) {
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		mockTokenStr, err := New(config.Get()).GenToken(mockUserID, mockTokenType, mockIat, mockExp)
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		mockTokenMeta := auth.TokenMetadata{
			UserID: mockUserID,
			Type:   mockTokenType,
			Token:  mockTokenStr,
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.AddCookie(&http.Cookie{
			Name:     "REFRESH_TOKEN",
			Value:    mockTokenMeta.Token,
			MaxAge:   3600,
			Path:     "/",
			Domain:   "/",
			Secure:   false,
			HttpOnly: true,
		})
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateRefreshToken(ctx)
		if assert.NoError(t, err) {
			assert.Equal(t, mockTokenMeta, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateRefreshToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "refresh token not found", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// wrong signing method
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"user_id": mockUserID,
			"type":    mockTokenType,
			"iat":     mockIat,
			"exp":     mockExp,
		})
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		tokenStr, err := token.SignedString(key)
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.AddCookie(&http.Cookie{
			Name:     "REFRESH_TOKEN",
			Value:    tokenStr,
			MaxAge:   3600,
			Path:     "/",
			Domain:   "/",
			Secure:   false,
			HttpOnly: true,
		})
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateRefreshToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid refresh token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// wrong secret
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": mockUserID,
			"type":    mockTokenType,
			"iat":     mockIat,
			"exp":     mockExp,
		})
		tokenStr, err := token.SignedString([]byte("some-other-secret"))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.AddCookie(&http.Cookie{
			Name:     "REFRESH_TOKEN",
			Value:    tokenStr,
			MaxAge:   3600,
			Path:     "/",
			Domain:   "/",
			Secure:   false,
			HttpOnly: true,
		})
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateRefreshToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid refresh token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// expired token
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": mockUserID,
			"type":    mockTokenType,
			"iat":     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			"exp":     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		})
		tokenStr, err := token.SignedString([]byte(config.Get().Jwt.Secret))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.AddCookie(&http.Cookie{
			Name:     "REFRESH_TOKEN",
			Value:    tokenStr,
			MaxAge:   3600,
			Path:     "/",
			Domain:   "/",
			Secure:   false,
			HttpOnly: true,
		})
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateRefreshToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid refresh token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})

	t.Run("it should return error", func(t *testing.T) {
		// no claims
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/some-url", strings.NewReader(""))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
		tokenStr, err := token.SignedString([]byte(config.Get().Jwt.Secret))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.AddCookie(&http.Cookie{
			Name:     "REFRESH_TOKEN",
			Value:    tokenStr,
			MaxAge:   3600,
			Path:     "/",
			Domain:   "/",
			Secure:   false,
			HttpOnly: true,
		})
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		tokenMeta, err := ss.GetAndValidateRefreshToken(ctx)
		if assert.Error(t, err) {
			assert.Equal(t, "invalid refresh token", err.Error())
			assert.Equal(t, auth.TokenMetadata{}, tokenMeta)
		}
	})
}
