package dto

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type PersonDto struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"nickname"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	Stacks    []string  `json:"stacks"`
}

type RequestCreatePersonDto struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"nickname"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	Stacks    []string  `json:"stacks"`
}

type ResponseGetPersonDto struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"nickname"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	Stacks    []string  `json:"stacks"`
}

type ResponseCreatePersonDto struct {
	ID uuid.UUID `json:"id"`
}

func (dto *PersonDto) IsEmpty() bool {
	return reflect.DeepEqual(&PersonDto{}, dto)
}

func (dto *RequestCreatePersonDto) GetSearchByTerm() string {
	return fmt.Sprintf("%s || %s || %s || %s || %s", dto.ID, dto.Nickname, dto.Name, dto.BirthDate, dto.GetConcatStaks())
}

func (dto *RequestCreatePersonDto) GetConcatStaks() string {
	return strings.Join(dto.Stacks, ",")
}

func (dto *ResponseGetPersonDto) IsEmpty() bool {
	return reflect.DeepEqual(&ResponseGetPersonDto{}, dto)
}
