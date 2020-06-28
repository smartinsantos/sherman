package exception

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewUnAuthorizedError(t *testing.T) {
	mockErrorMessage := "some-error-message"
	err := NewUnAuthorizedError(mockErrorMessage)
	assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&UnAuthorizedError{}))
	assert.Equal(t, mockErrorMessage, err.Error())
}
