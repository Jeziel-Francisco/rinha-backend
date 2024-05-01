package service

import (
	"docker-example/src/application/ports/out"
	"docker-example/src/commons/errors"
)

func NewCheckExistsNicknameService(personAdapter out.PersonDatabaseAdapter) Service {
	return &checkExistsnickName{
		personAdapter: personAdapter,
	}
}

type checkExistsnickName struct {
	personAdapter out.PersonDatabaseAdapter
}

func (service *checkExistsnickName) Execute(args ...interface{}) (interface{}, errors.CommonError) {
	if len(args) == 0 {
		return nil, errors.NewInvaliParamApiError("args")
	}

	nickname, ok := args[0].(string)
	if !ok {
		return nil, errors.NewInvaliParamApiError("nickname")
	}

	person, err := service.personAdapter.GetPersonByNickname(nickname)
	if err != nil {
		return nil, err
	}
	if person != nil {
		return nil, errors.NewPersonAlreadyExists()
	}

	return nil, nil
}
