package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"sherman/src/app/config"
	"testing"
	"time"
)

func TestValidateHash(t *testing.T) {
	mockPassword := "some-password"
	actualHash, err := New().Hash(mockPassword)
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

	ss := New()
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

	tokenStr, err := New().GenToken(mockUserID, mockTokenType, mockIat, mockExp)
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
