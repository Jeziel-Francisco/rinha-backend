package person

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/adapter/mapper"
	"docker-example/src/ports/out/drive/database/person"
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

func (adapter *personDatabaseAdapter) Create(person *entities.Person) (*entities.Person, errors.CommonError) {
	result, err := adapter.personDatabase.Create(mapper.FromPersonToRequestCreatePersonDto(person))
	if err != nil {
		return nil, err
	}
	return &entities.Person{
		ID: result.ID.String(),
	}, nil
}
