package person

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/commons/errors"

	"github.com/stretchr/testify/mock"
)

func NewPersonDatabaseAdapterMock() *personDatabaseAdapterMock {
	return &personDatabaseAdapterMock{}
}

type personDatabaseAdapterMock struct {
	mock.Mock
}

func (adapter *personDatabaseAdapterMock) GetPersonByNickname(nickname string) (*entities.Person, errors.CommonError) {
	args := adapter.Called(nickname)
	responseSuccess, responseError := args[0], args[1]

	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(*entities.Person), nil
}

func (adapter *personDatabaseAdapterMock) Create(person *entities.Person) (*entities.Person, errors.CommonError) {
	args := adapter.Called(person)
	responseSuccess, responseError := args[0], args[1]

	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(*entities.Person), nil
}
