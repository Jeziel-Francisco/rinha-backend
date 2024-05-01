package usecase

import (
	"github.com/stretchr/testify/mock"
)

type personCreateUseCaseMock struct {
	mock.Mock
}

func NewPersonCreateUseCaseMock() *personCreateUseCaseMock {
	return &personCreateUseCaseMock{}
}

func (useCaseMock *personCreateUseCaseMock) Execute(intention interface{}) (err error) {
	args := useCaseMock.Called(intention)
	responseError := args[0]
	if responseError != nil {
		return responseError.(error)
	}
	return nil
}
