package errors

import (
	"fmt"
	"net/http"
)

type BusinessError struct {
	status  int
	message string
	code    string
}

func (err *BusinessError) Status() int {
	return err.status
}

func (err *BusinessError) Message() string {
	return err.message
}

func (err *BusinessError) Code() string {
	return err.code
}

func (err *BusinessError) Body() Error {
	return Error{
		Code:    err.code,
		Message: err.message,
	}
}

func NewBusinessError(status int, message string, code string) *BusinessError {
	return &BusinessError{
		status:  status,
		message: message,
		code:    code,
	}
}

func NewRequiredFieldError(fielName string) *BusinessError {
	return NewBusinessError(http.StatusBadRequest, fmt.Sprintf("%s is required!", fielName), "required_field")
}

func NewInvaliFieldError(fielName string) *BusinessError {
	return NewBusinessError(http.StatusBadRequest, fmt.Sprintf("%s is invalid!", fielName), "invalid_field")
}

func NewPersonAlreadyExists() *BusinessError {
	return NewBusinessError(http.StatusBadRequest, "person already exists", "person_already_exists")
}

func NewPersonNotFound() *BusinessError {
	return NewBusinessError(http.StatusNotFound, "person not found", "person_not_found")
}
