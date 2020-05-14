package response

import "net/http"

// Response struct for response shape
type Response struct {
	Status int
	Error string
	Errors map[string]string
	Data map[string]interface{}
}

func NewResponse() Response {
	return Response{
		Status: http.StatusInternalServerError,
	}
}

// GetStatus gets the status of the response struct instance defaults to 500
func (res *Response) GetStatus() int {
	return res.Status
}

// GetBody gets the body of a response struct
func (res *Response) GetBody() map[string]interface{} {
	response := make(map[string]interface{})

	if len(res.Error) > 0 {
		response["error"] = res.Error
	}
	if len(res.Errors) > 0 {
		response["errors"] = res.Errors
	}

	response["data"] = res.Data
	return response
}

func (res *Response) SetInternalServerError() {
	res.Status = http.StatusInternalServerError
	res.Error = "internal server error"
	res.Errors = nil
	res.Data = nil
}

func (res *Response) SetError(status int, error string) {
	res.Status = status
	res.Error = error
	res.Errors = nil
	res.Data = nil
}

func (res *Response) SetErrors(status int, errors map[string]string) {
	res.Status = status
	res.Errors = errors
	res.Error = ""
	res.Data = nil
}

func (res *Response) SetData(status int, data map[string]interface{}) {
	res.Status = status
	res.Data = data
	res.Error = ""
	res.Errors = nil
}