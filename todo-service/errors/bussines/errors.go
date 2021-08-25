package bussines

import (
	"fmt"
	"todo-service/errors"
)

const (
	Unknown          int = 0
	BussinesErrType1 int = iota + 1
	BussinesErrType2
	BussinesErrType3
)

type BadFieldForEntity struct {
	myerrors.BaseError
	EntityType string
	EntityID   int
}

func (e *BadFieldForEntity) Error() string {
	return fmt.Sprintf("badfieldforentity code %v message %v id %v type %v base: %v", e.Code, e.Message, e.EntityID, e.EntityType, e.Err)
}

//func (e *BadFieldForEntity) Unwrap() error { return e.Err }

func NewBadFieldForEntity(baseError myerrors.BaseError, entityType string, entityID int) *BadFieldForEntity {
	return &BadFieldForEntity{BaseError: baseError, EntityType: entityType, EntityID: entityID}
}
