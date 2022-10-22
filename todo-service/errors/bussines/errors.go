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

// Unwrap to allow use As for assertion type
func (e *BadFieldForEntity) Unwrap() error { return e.Err }

// NewBadFieldForEntity allow creates a new BadFieldForEntity
func NewBadFieldForEntity(baseError myerrors.BaseError, entityType string, entityID int) *BadFieldForEntity {
	return &BadFieldForEntity{BaseError: baseError, EntityType: entityType, EntityID: entityID}
}
