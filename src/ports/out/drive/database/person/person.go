package person

import (
	"context"
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/database/dto"
	"docker-example/src/ports/out/infraestructure"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	scriptGetPersonByNickname = "select id, nickname, name, birthdate, stack from public.person where nickname = $1"
	scriptGetPersonByID       = "select id, nickname, name, birthdate, stack from public.person where id = $1"
	scriptGetPersonByTerm     = "select id, nickname, name, birthdate, stack from public.person where search LIKE '%' || $1 || '%'"
	scriptCountPerson         = "select COUNT(1) from public.person"
	scriptInsertPerson        = "INSERT INTO public.person (id, nickname, name, birthdate, stack, search) VALUES ($1, $2, $3, $4, $5, $6);"
)

func NewPersonDatabase(database infraestructure.Database) PersonDatabase {
	return &personDatabase{
		database: database,
	}
}

type PersonDatabase interface {
	GetPersonByNickname(nickname string) (*dto.ResponseGetPersonDto, errors.CommonError)
	GetPersonByID(ID uuid.UUID) (*dto.ResponseGetPersonDto, errors.CommonError)
	GetPersonByTerm(term string) ([]*dto.ResponseGetPersonDto, errors.CommonError)
	Create(person *dto.RequestCreatePersonDto) (*dto.ResponseCreatePersonDto, errors.CommonError)
	CountPerson() (*uint64, errors.CommonError)
}

type personDatabase struct {
	database infraestructure.Database
}

func (client *personDatabase) GetPersonByNickname(nickname string) (*dto.ResponseGetPersonDto, errors.CommonError) {
	var resultNickname, name, stacks string
	var birthDate time.Time
	var ID uuid.UUID

	err := client.database.GetPoolConnection().QueryRow(context.Background(), scriptGetPersonByNickname, nickname).Scan(&ID, &resultNickname, &name, &birthDate, &stacks)
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

func (client *personDatabase) GetPersonByID(ID uuid.UUID) (*dto.ResponseGetPersonDto, errors.CommonError) {
	var nickname, name, stacks string
	var birthDate time.Time

	err := client.database.GetPoolConnection().QueryRow(context.Background(), scriptGetPersonByID, ID).Scan(&ID, &nickname, &name, &birthDate, &stacks)
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

func (client *personDatabase) GetPersonByTerm(term string) ([]*dto.ResponseGetPersonDto, errors.CommonError) {
	var nickname, name, stacks string
	var birthDate time.Time
	var ID uuid.UUID

	rows, err := client.database.GetPoolConnection().Query(context.Background(), scriptGetPersonByTerm, term)
	if client.database.HasError(err) {
		return nil, errors.NewClientErrorByError(err)
	}
	if client.database.HasEmptyData(err) {
		return nil, nil
	}
	results, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*dto.ResponseGetPersonDto, error) {
		row.Scan(&ID, &nickname, &name, &birthDate, &stacks)
		return &dto.ResponseGetPersonDto{
			ID:        ID,
			Nickname:  nickname,
			Name:      name,
			BirthDate: birthDate.String(),
			Stacks:    strings.Split(stacks, ","),
		}, nil
	})
	if err != nil {
		return nil, errors.NewClientErrorByError(err)
	}
	return results, nil
}

func (client *personDatabase) CountPerson() (*uint64, errors.CommonError) {
	var quantityPerson uint64

	err := client.database.GetPoolConnection().QueryRow(context.Background(), scriptCountPerson).Scan(&quantityPerson)
	if client.database.HasError(err) {
		return nil, errors.NewClientErrorByError(err)
	}
	if client.database.HasEmptyData(err) {
		return nil, nil
	}
	return &quantityPerson, nil
}

func (client *personDatabase) Create(person *dto.RequestCreatePersonDto) (*dto.ResponseCreatePersonDto, errors.CommonError) {
	person.ID = uuid.New()

	_, err := client.database.GetPoolConnection().Exec(context.Background(), scriptInsertPerson, person.ID, person.Nickname, person.Name, person.BirthDate, person.GetConcatStaks(), person.GetSearchByTerm())
	if err != nil {
		return nil, errors.NewClientErrorByError(err)
	}
	return &dto.ResponseCreatePersonDto{ID: person.ID}, nil
}
