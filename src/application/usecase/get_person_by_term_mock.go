package usecase

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

type getPersonByTermUseCaseMock struct {
	mock.Mock
}

func NewGetPersonByTermUseCaseMock() *getPersonByTermUseCaseMock {
	return &getPersonByTermUseCaseMock{}
}

func (useCaseMock *getPersonByTermUseCaseMock) Execute(intention interface{}) (err errors.CommonError) {
	args := useCaseMock.Called(intention)
	responseError := args[0]
	if responseError != nil {
		return responseError.(errors.CommonError)
	}
	return nil
}
