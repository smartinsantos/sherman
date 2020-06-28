package response

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

var (
	mockStatus = 0
	mockError  = "some-error"
	mockErrors = map[string]string{"some-error": "some-error", "another-error": "another-error"}
	mockData   = D{"some-data": "some-data"}
)

func TestResponse(t *testing.T) {
	response := NewResponse()
	assert.NotEmpty(t, response)
	assert.Equal(t, reflect.TypeOf(response), reflect.TypeOf(&Response{}))
	assert.Equal(t, http.StatusInternalServerError, response.Status)
	assert.Equal(t, internalServerError, response.Error)
	assert.Nil(t, response.Errors)
	assert.Nil(t, response.Data)
}

func TestGetStatus(t *testing.T) {
	response := NewResponse()
	assert.Equal(t, http.StatusInternalServerError, response.GetStatus())
	response.Status = mockStatus
	assert.Equal(t, mockStatus, response.GetStatus())
}

func TestGetBody(t *testing.T) {
	response := NewResponse()
	body := response.GetBody()
	assert.Equal(t, internalServerError, body["error"])
	assert.Nil(t, body["data"])
	response.Error = mockError
	response.Errors = mockErrors
	response.Data = mockData
	body = response.GetBody()
	assert.Equal(t, mockError, body["error"])
	assert.Equal(t, mockErrors, body["errors"])
	assert.Equal(t, mockData, body["data"])
}

func TestSetInternalServerError(t *testing.T) {
	response := NewResponse()
	response.Status = mockStatus
	response.Error = mockError
	response.Errors = mockErrors
	response.Data = mockData

	response.SetInternalServerError()
	assert.Equal(t, http.StatusInternalServerError, response.Status)
	assert.Equal(t, internalServerError, response.Error)
	assert.Nil(t, response.Errors)
	assert.Nil(t, response.Data)
}

func TestSetError(t *testing.T) {
	response := NewResponse()
	response.SetError(mockStatus, mockError)
	assert.Equal(t, mockStatus, response.Status)
	assert.Equal(t, mockError, response.Error)
}

func TestSetErrors(t *testing.T) {
	response := NewResponse()
	response.SetErrors(mockStatus, mockErrors)
	assert.Equal(t, mockStatus, response.Status)
	assert.Equal(t, mockErrors, response.Errors)
}

func TestSetData(t *testing.T) {
	response := NewResponse()
	response.SetData(mockStatus, mockData)
	assert.Equal(t, mockStatus, response.Status)
	assert.Equal(t, mockData, response.Data)
}
