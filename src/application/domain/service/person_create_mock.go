package service

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

func NewPersonCreateServiceMock() *personCreateMock {
	return &personCreateMock{}
}

type personCreateMock struct {
	mock.Mock
}

func (serviceMock *personCreateMock) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	mockArgs := serviceMock.Called(args)
	responseSuccess, responseError := mockArgs[0], mockArgs[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(string), nil
}
