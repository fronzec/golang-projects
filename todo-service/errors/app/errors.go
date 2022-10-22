package app

import (
	"fmt"
	"todo-service/errors"
)

type BadConfigErr struct {
	myerrors.BaseError
}

func NewBadConfigErr(baseError myerrors.BaseError) *BadConfigErr {
	return &BadConfigErr{BaseError: baseError}
}

func (e *BadConfigErr) Error() string {
	return fmt.Sprintf("BadConfigErr code %v message %v base: %v", e.Code, e.Message, e.Err)
}

// Is to compare one error values vs other
func (e *BadConfigErr) Is(target error) bool {
	t, ok := target.(*BadConfigErr)
	if !ok {
		return false
	}
	return (e.Message == t.Message || t.Message == "") &&
		(e.Code == t.Code || t.Code == 0)
}
