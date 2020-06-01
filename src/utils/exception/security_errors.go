package exception

// UnAuthorizedError struct error type that should be used to indicate that the error is caused by the nonexistence of the requested resource.
type UnAuthorizedError struct {
	msg string
}

// NewUnAuthorizedError is the ErrValidation constructor.
func NewUnAuthorizedError(msg string) *UnAuthorizedError {
	return &UnAuthorizedError{msg: msg}
}

// Error returns the error message.
func (err *UnAuthorizedError) Error() string {
	return err.msg
}
