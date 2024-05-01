package errors

import "net/http"

type ApiError struct {
	status  int
	message string
	code    string
}

func (err *ApiError) Status() int {
	return err.status
}

func (err *ApiError) Message() string {
	return err.message
}

func (err *ApiError) Code() string {
	return err.code
}

func (err *ApiError) Body() Error {
	return Error{
		Code:    err.code,
		Message: err.message,
	}
}

func NewApiError(status int, message string, code string) *ApiError {
	return &ApiError{
		status:  status,
		message: message,
		code:    code,
	}
}

func NewUnmarshalApiError(causes string) *ApiError {
	return NewApiError(http.StatusBadRequest, causes, "unmarshal_error")
}

func NewInvaliParamApiError(causes string) *ApiError {
	return NewApiError(http.StatusBadRequest, causes, "invalid_param")
}
