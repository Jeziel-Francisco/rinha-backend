package person

import (
	"context"
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/database/dto"
	"docker-example/src/ports/out/infraestructure"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	scriptGetPersonByID = "select id, nickname, name, birthdate, stack from public.person where nickname = $1"
	scriptInsertPerson  = "INSERT INTO public.person (id, nickname, name, birthdate, stack, search) VALUES ($1, $2, $3, $4, $5, $6);"
)

func NewPersonDatabase(database infraestructure.Database) PersonDatabase {
	return &personDatabase{
		database: database,
	}
}

type PersonDatabase interface {
	GetPersonByNickname(nickname string) (*dto.ResponseGetPersonDto, errors.CommonError)
	Create(person *dto.RequestCreatePersonDto) (*dto.ResponseCreatePersonDto, errors.CommonError)
}

type personDatabase struct {
	database infraestructure.Database
}

func (client *personDatabase) GetPersonByNickname(nickname string) (*dto.ResponseGetPersonDto, errors.CommonError) {
	var resultNickname, name, stacks string
	var birthDate time.Time
	var ID uuid.UUID

	err := client.database.GetPoolConnection().QueryRow(context.Background(), scriptGetPersonByID, nickname).Scan(&ID, &resultNickname, &name, &birthDate, &stacks)
	if client.database.HasError(err) {
		return nil, errors.NewClientErrorByError(err)
	}
	if client.database.HasEmptyData(err) {
		return nil, nil
	}
	return &dto.ResponseGetPersonDto{
		ID:        ID,
		Nickname:  nickname,
		Name:      name,
		BirthDate: birthDate.String(),
		Stacks:    strings.Split(stacks, ","),
	}, nil
}

func (client *personDatabase) Create(person *dto.RequestCreatePersonDto) (*dto.ResponseCreatePersonDto, errors.CommonError) {
	person.ID = uuid.New()

	_, err := client.database.GetPoolConnection().Exec(context.Background(), scriptInsertPerson, person.ID, person.Nickname, person.Name, person.BirthDate, person.GetConcatStaks(), person.GetSearchByTerm())
	if err != nil {
		return nil, errors.NewClientErrorByError(err)
	}
	return &dto.ResponseCreatePersonDto{ID: person.ID}, nil
}