package dto

import (
	"errors"
	"fmt"
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

type Person struct {
	Nickname  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	Stacks    []string `json:"stack"`
}

func (person *Person) Validate() error {
	if len(person.Nickname) == minimalLenNickname || len(person.Nickname) > maxLenNickname {
		return errors.New("Apelido inválido")
	}
	if len(person.Name) == minimalLenName || len(person.Name) > maxLenName {
		return errors.New("Nome inválido")
	}

	_, err := time.Parse("2006-01-02", person.BirthDate)
	if err != nil {
		return errors.New("Data de nascimento inválida")
	}

	for _, stack := range person.Stacks {
		if len(stack) == minimalLenStackItem || len(stack) > maxLenStackItem {
			return errors.New(fmt.Sprintf("Stack item inválido, %s", stack))
		}
	}
	return nil
}
