package myerrors

import "strconv"

type BaseError struct {
	Err     error
	Message string
	Code    int
}

func NewBaseError(err error, message string, code int) *BaseError {
	return &BaseError{Err: err, Message: message, Code: code}
}

func (e *BaseError) Error() string {
	return strconv.Itoa(e.Code) + " " + e.Message + ": " + e.Err.Error()
}
