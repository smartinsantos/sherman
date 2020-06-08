package response

import (
	"net/http"
)

var internalServerError = "internal server error"

type (
	// D Response Data type
	D map[string]interface{}

	// Response Response.Response struct definition
	Response struct {
		Status int
		Error  string
		Errors map[string]string
		Data   D
	}
)

// NewResponse Response constructor defaults to { Status: 500, error: "internal server error" }
func NewResponse() Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  internalServerError,
	}
}

// GetStatus returns the status of the Response
func (res *Response) GetStatus() int {
	return res.Status
}

// GetBody returns the body of the Response contains status key, and one of the following keys: error, errors, data
func (res *Response) GetBody() map[string]interface{} {
	body := make(map[string]interface{})

	if len(res.Error) > 0 {
		body["error"] = res.Error
	}
	if len(res.Errors) > 0 {
		body["errors"] = res.Errors
	}

	body["data"] = res.Data
	return body
}

// SetInternalServerError sets the Response to internal server error { Status: 500, error: "internal server error" }
func (res *Response) SetInternalServerError() {
	res.Status = http.StatusInternalServerError
	res.Error = internalServerError
	res.Errors = nil
	res.Data = nil
}

// SetError sets with an error { Status: [status], Error: [error], Errors: nil, Data: nil }
func (res *Response) SetError(status int, err string) {
	res.Status = status
	res.Error = err
	res.Errors = nil
	res.Data = nil
}

// SetErrors sets a Response with multiple errors { Status: [status], Errors: [errors], Error: "", Data: nil }
func (res *Response) SetErrors(status int, errs map[string]string) {
	res.Status = status
	res.Errors = errs
	res.Error = ""
	res.Data = nil
}

// SetData sets a Response with data { Status: [status], Data: [data], Errors: nil, Error: "" }
func (res *Response) SetData(status int, data map[string]interface{}) {
	res.Status = status
	res.Data = data
	res.Error = ""
	res.Errors = nil
}
