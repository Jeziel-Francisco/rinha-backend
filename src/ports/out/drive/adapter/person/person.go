package person

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/adapter/mapper"
	"docker-example/src/ports/out/drive/database/person"

	"github.com/google/uuid"
)

func NewPersonDatabaseAdapter(personDatabase person.PersonDatabase) out.PersonDatabaseAdapter {
	return &personDatabaseAdapter{
		personDatabase: personDatabase,
	}
}

type personDatabaseAdapter struct {
	personDatabase person.PersonDatabase
}

func (adapter *personDatabaseAdapter) GetPersonByNickname(nickname string) (*entities.Person, errors.CommonError) {
	result, err := adapter.personDatabase.GetPersonByNickname(nickname)
	if err != nil {
		return nil, err
	}
	return mapper.FromResponseGetPersonDtoToPerson(result), nil
}

func (adapter *personDatabaseAdapter) GetPersonByID(ID string) (*entities.Person, errors.CommonError) {
	newID, _ := uuid.Parse(ID)
	result, err := adapter.personDatabase.GetPersonByID(newID)
	if err != nil {
		return nil, err
	}
	return mapper.FromResponseGetPersonDtoToPerson(result), nil
}

func (adapter *personDatabaseAdapter) GetPersonByTerm(term string) ([]*entities.Person, errors.CommonError) {
	result, err := adapter.personDatabase.GetPersonByTerm(term)
	if err != nil {
		return nil, err
	}
	return mapper.FromListResponseGetPersonDtoToListPerson(result), nil
}

func (adapter *personDatabaseAdapter) CountPerson() (*uint64, errors.CommonError) {
	result, err := adapter.personDatabase.CountPerson()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (adapter *personDatabaseAdapter) Create(person *entities.Person) (*entities.Person, errors.CommonError) {
	result, err := adapter.personDatabase.Create(mapper.FromPersonToRequestCreatePersonDto(person))
	if err != nil {
		return nil, err
	}
	return &entities.Person{
		ID: result.ID.String(),
	}, nil
}
