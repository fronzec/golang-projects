package types

type ErrorResponse struct {
	Message  string
	HTTPCode int
}

func (e *ErrorResponse) Error() string {
	return e.Message
}