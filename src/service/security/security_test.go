package security

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
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

}
