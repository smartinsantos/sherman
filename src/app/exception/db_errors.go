package exception

// NotFoundError struct error type that should be used to indicate that the error is caused by the nonexistence of the requested resource.
type NotFoundError struct {
	msg string
}
// NewNotFoundError is the ErrValidation constructor.
func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError { msg: msg }
}
// Error returns the error message.
func (err *NotFoundError) Error() string {
	return err.msg
}

// DuplicateEntryError struct error type that should be used to indicate that the error is caused by the nonexistence of the requested resource.
type DuplicateEntryError struct {
	msg string
}
// NewDuplicateEntryError is the ErrValidation constructor.
func NewDuplicateEntryError(msg string) *DuplicateEntryError {
	return &DuplicateEntryError { msg: msg }
}
// Error returns the error message.
func (err *DuplicateEntryError) Error() string {
	return err.msg
}
