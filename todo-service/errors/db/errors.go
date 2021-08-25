package db

import (
	"fmt"
	"todo-service/errors"
)

type RequiredRecordNotFound struct {
	myerrors.BaseError
	DBCode int
}

func (e *RequiredRecordNotFound) Error() string{
	return fmt.Sprintf("requiredrecordnotfound code %v message %v dbcode %v base: %v",
		e.Code, e.Message, e.DBCode, e.Err)
}

func NewRequiredRecordNotFound(baseError myerrors.BaseError, DBCode int) *RequiredRecordNotFound {
	return &RequiredRecordNotFound{BaseError: baseError, DBCode: DBCode}
}