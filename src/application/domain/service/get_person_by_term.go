package service

import (
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
)

func NewGetPersonByTermService(personAdapter out.PersonDatabaseAdapter) Service {
	return &getPersonByTerm{
		personAdapter: personAdapter,
	}
}

type getPersonByTerm struct {
	personAdapter out.PersonDatabaseAdapter
}

func (service *getPersonByTerm) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	if len(args) == 0 {
		return nil, errors.NewInvaliParamApiError("args")
	}

	term, ok := args[0].(string)
	if !ok {
		return nil, errors.NewInvaliParamApiError("id")
	}

	people, err := service.personAdapter.GetPersonByTerm(term)
	if err != nil {
		return nil, err
	}
	return people, nil
}
