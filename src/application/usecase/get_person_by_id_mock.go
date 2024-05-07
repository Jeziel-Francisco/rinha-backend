package usecase

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

type getPersonByIDUseCaseMock struct {
	mock.Mock
}

func NewGetPersonByIDUseCaseMock() *getPersonByIDUseCaseMock {
	return &getPersonByIDUseCaseMock{}
}

func (useCaseMock *getPersonByIDUseCaseMock) Execute(intention interface{}) (err errors.CommonError) {
	args := useCaseMock.Called(intention)
	responseError := args[0]
	if responseError != nil {
		return responseError.(errors.CommonError)
	}
	return nil
}
