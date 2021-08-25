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

func (b *BadConfigErr) Error() string {
	return fmt.Sprintf("BadConfigErr code %v message %v base: %v", b.Code, b.Message, b.Err)
}