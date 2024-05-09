package service

import (
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
)

func NewCountPersonService(personAdapter out.PersonDatabaseAdapter) Service {
	return &countPersonService{
		personAdapter: personAdapter,
	}
}

type countPersonService struct {
	personAdapter out.PersonDatabaseAdapter
}

func (service *countPersonService) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	quantityPerson, err := service.personAdapter.CountPerson()
	if err != nil {
		return nil, err
	}
	return quantityPerson, nil
}
