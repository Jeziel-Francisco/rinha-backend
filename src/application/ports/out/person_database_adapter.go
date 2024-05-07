package out

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/commons/errors"
)

type PersonDatabaseAdapter interface {
	GetPersonByNickname(nickName string) (*entities.Person, errors.CommonError)
	GetPersonByID(ID string) (*entities.Person, errors.CommonError)
	GetPersonByTerm(term string) ([]*entities.Person, errors.CommonError)
	Create(person *entities.Person) (*entities.Person, errors.CommonError)
}
