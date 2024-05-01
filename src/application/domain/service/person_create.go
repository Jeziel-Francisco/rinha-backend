package service

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
)

func NewPersonCreateService(personAdapter out.PersonDatabaseAdapter) Service {
	return &personCreate{
		personAdapter: personAdapter,
	}
}

type personCreate struct {
	personAdapter out.PersonDatabaseAdapter
}

func (service *personCreate) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	if len(args) == 0 {
		return nil, errors.NewInvaliParamApiError("args")
	}

	person, ok := args[0].(entities.Person)
	if !ok {
		return nil, errors.NewInvaliParamApiError("person")
	}

	newPerson, err := service.personAdapter.Create(&person)
	if err != nil {
		return nil, err
	}

	return newPerson.ID, nil
}
