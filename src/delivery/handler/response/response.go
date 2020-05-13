package response

import "net/http"

// Response struct for response shape
type Response struct {
	Status int
	Error string
	Errors map[string]string
	Data map[string]interface{}
}

// GetStatus gets the status of the response struct instance defaults to 500
func (res *Response) GetStatus() int {
	if res.Status == 0 {
		return http.StatusInternalServerError
	}
	return res.Status
}

// GetBody gets the body of a response struct instance defaults to { error: "internal server error" }
func (res *Response) GetBody() map[string]interface{} {
	response := make(map[string]interface{})

	if res.Status == 0 {
		response["error"] = "internal server error"
		return response
	}
	if len(res.Error) > 0 {
		response["error"] = res.Error
	}
	if len(res.Errors) > 0 {
		response["errors"] = res.Errors
	}

	response["data"] = res.Data
	return response
}
