package exception

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewNotFoundError(t *testing.T) {
	mockErrorMessage := "some-error-message"
	err := NewNotFoundError(mockErrorMessage)
	assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&NotFoundError{}))
	assert.Equal(t, mockErrorMessage, err.Error())
}

func TestDuplicateEntryError(t *testing.T) {
	mockErrorMessage := "some-error-message"
	err := NewDuplicateEntryError(mockErrorMessage)
	assert.Equal(t, reflect.TypeOf(err), reflect.TypeOf(&DuplicateEntryError{}))
	assert.Equal(t, mockErrorMessage, err.Error())
}
