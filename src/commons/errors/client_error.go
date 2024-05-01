package errors

import "net/http"

type ClientError struct {
	status  int
	message string
	code    string
}

func (err *ClientError) Status() int {
	return err.status
}

func (err *ClientError) Message() string {
	return err.message
}

func (err *ClientError) Code() string {
	return err.code
}

func (err *ClientError) Body() Error {
	return Error{
		Code:    err.code,
		Message: err.message,
	}
}

func NewClientError(status int, message string, code string) *ClientError {
	return &ClientError{
		status:  status,
		message: message,
		code:    code,
	}
}

func NewClientErrorByError(err error) *ClientError {
	return &ClientError{
		status:  http.StatusInternalServerError,
		message: err.Error(),
	}
}
