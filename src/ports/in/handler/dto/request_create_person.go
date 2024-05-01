package dto

import (
	"docker-example/src/commons/errors"
	"time"
)

const (
	minimalLenNickname  = 0
	maxLenNickname      = 32
	minimalLenName      = 0
	maxLenName          = 100
	minimalLenStackItem = 0
	maxLenStackItem     = 32
)

type RequestCreatePersonDto struct {
	Nickname  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	Stacks    []string `json:"stack"`
}

func (person *RequestCreatePersonDto) Validate() errors.CommonError {
	if len(person.Nickname) == minimalLenNickname || len(person.Nickname) > maxLenNickname {
		return errors.NewInvaliFieldError("apelido")
	}
	if len(person.Name) == minimalLenName || len(person.Name) > maxLenName {
		return errors.NewInvaliFieldError("nome")
	}

	_, err := time.Parse("2006-01-02", person.BirthDate)
	if err != nil {
		return errors.NewInvaliFieldError("nascimento")
	}

	for _, stack := range person.Stacks {
		if len(stack) == minimalLenStackItem || len(stack) > maxLenStackItem {
			return errors.NewInvaliFieldError(stack)
		}
	}
	return nil
}
