package errors

type CommonError interface {
	Status() int
	Message() string
	Code() string
	Body() Error
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func NewError(commonError CommonError) *Error {
	return &Error{
		Message: commonError.Message(),
		Code:    commonError.Code(),
	}
}
