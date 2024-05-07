package service

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

func NewGetPersonByTermServiceMock() *getPersonByTermServiceMock {
	return &getPersonByTermServiceMock{}
}

type getPersonByTermServiceMock struct {
	mock.Mock
}

func (serviceMock *getPersonByTermServiceMock) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	mockArgs := serviceMock.Called(args)
	responseSuccess, responseError := mockArgs[0], mockArgs[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.([]*entities.Person), nil
}
