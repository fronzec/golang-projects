package http

import (
	"fmt"
	"todo-service/errors"
)

type BadRequest struct {
	myerrors.BaseError
	URL string
	ResponsePayload string
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("BadRequest code %v message %v url %v response %v base: %v",
		e.Code, e.Message, e.URL, e.ResponsePayload, e.Err)
}

func NewBadRequest(baseError myerrors.BaseError, responsePayload string, URL string) *BadRequest {
	return &BadRequest{BaseError: baseError, ResponsePayload: responsePayload, URL: URL}
}

type AccessDenied struct {
	myerrors.BaseError
	URL string
}

func (e *AccessDenied) Error() string {
	return fmt.Sprintf("AccessDenied code %v message %v url %v base: %v",
		e.Code, e.Message, e.URL, e.Err)
}


func NewAccessDenied(baseError myerrors.BaseError, URL string) *AccessDenied {
	return &AccessDenied{BaseError: baseError, URL: URL}
}

type AccessUnauthorized struct {
	myerrors.BaseError
	URL string
}

func NewAccessUnauthorized(baseError myerrors.BaseError, URL string) *AccessUnauthorized {
	return &AccessUnauthorized{BaseError: baseError, URL: URL}
}

func (e *AccessUnauthorized) Error() string {
	return fmt.Sprintf("AccessUnauthorized code %v message %v url %v base: %v",
		e.Code, e.Message, e.URL, e.Err)
}

type ConnTimeoutErr struct {
	myerrors.BaseError
	URL string
}

type ResponseTimeoutErr struct {
	myerrors.BaseError
	URL string
}

type ClientError struct {
	myerrors.BaseError
	URL string
	HTTPCode int
}

type ServerError struct {
	myerrors.BaseError
	URL string
	HTTPCode int
}