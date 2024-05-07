package person

import (
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/database/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewPersonDatabaseMock() *personDatabaseMock {
	return &personDatabaseMock{}
}

type personDatabaseMock struct {
	mock.Mock
}

func (clientMock *personDatabaseMock) GetPersonByNickname(nickname string) (*dto.ResponseGetPersonDto, errors.CommonError) {
	args := clientMock.Called(nickname)
	responseSuccess, responseError := args[0], args[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(*dto.ResponseGetPersonDto), nil
}

func (clientMock *personDatabaseMock) GetPersonByID(ID uuid.UUID) (*dto.ResponseGetPersonDto, errors.CommonError) {
	args := clientMock.Called(ID)
	responseSuccess, responseError := args[0], args[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(*dto.ResponseGetPersonDto), nil
}

func (clientMock *personDatabaseMock) GetPersonByTerm(term string) ([]*dto.ResponseGetPersonDto, errors.CommonError) {
	args := clientMock.Called(term)
	responseSuccess, responseError := args[0], args[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.([]*dto.ResponseGetPersonDto), nil
}

func (clientMock *personDatabaseMock) Create(person *dto.RequestCreatePersonDto) (*dto.ResponseCreatePersonDto, errors.CommonError) {
	args := clientMock.Called(person)
	responseSuccess, responseError := args[0], args[1]
	if responseError != nil {
		return nil, responseError.(errors.CommonError)
	}
	return responseSuccess.(*dto.ResponseCreatePersonDto), nil
}
