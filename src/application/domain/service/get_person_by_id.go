package service

import (
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
)

func NewGetPersonByIDService(personAdapter out.PersonDatabaseAdapter) Service {
	return &getPersonByID{
		personAdapter: personAdapter,
	}
}

type getPersonByID struct {
	personAdapter out.PersonDatabaseAdapter
}

func (service *getPersonByID) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	if len(args) == 0 {
		return nil, errors.NewInvaliParamApiError("args")
	}

	personID, ok := args[0].(string)
	if !ok {
		return nil, errors.NewInvaliParamApiError("id")
	}

	person, err := service.personAdapter.GetPersonByID(personID)
	if err != nil {
		return nil, err
	}
	if person == nil {
		return nil, errors.NewPersonNotFound()
	}
	return person, nil
}
