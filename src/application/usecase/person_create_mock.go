package usecase

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

type personCreateUseCaseMock struct {
	mock.Mock
}

func NewPersonCreateUseCaseMock() *personCreateUseCaseMock {
	return &personCreateUseCaseMock{}
}

func (useCaseMock *personCreateUseCaseMock) Execute(intention interface{}) (err errors.CommonError) {
	args := useCaseMock.Called(intention)
	responseError := args[0]
	if responseError != nil {
		return responseError.(errors.CommonError)
	}
	return nil
}
