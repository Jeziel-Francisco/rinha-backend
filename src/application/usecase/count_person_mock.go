package usecase

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

type countPersonUseCaseMock struct {
	mock.Mock
}

func NewCountPersonUseCaseMock() *countPersonUseCaseMock {
	return &countPersonUseCaseMock{}
}

func (useCaseMock *countPersonUseCaseMock) Execute(intention interface{}) (err errors.CommonError) {
	args := useCaseMock.Called(intention)
	responseError := args[0]
	if responseError != nil {
		return responseError.(errors.CommonError)
	}
	return nil
}
