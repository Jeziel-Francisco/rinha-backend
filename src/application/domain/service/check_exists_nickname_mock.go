package service

import (
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

func NewCheckExistsNicknameServiceMock() *checkExistsnickNameMock {
	return &checkExistsnickNameMock{}
}

type checkExistsnickNameMock struct {
	mock.Mock
}

func (serviceMock *checkExistsnickNameMock) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	mockArgs := serviceMock.Called(args)
	_, responseError := mockArgs[0], mockArgs[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return nil, nil
}
