package service

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

func NewCountPersonServiceMock() *countPersonServiceMock {
	return &countPersonServiceMock{}
}

type countPersonServiceMock struct {
	mock.Mock
}

func (serviceMock *countPersonServiceMock) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	mockArgs := serviceMock.Called()
	responseSuccess, responseError := mockArgs[0], mockArgs[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(*uint64), nil
}
