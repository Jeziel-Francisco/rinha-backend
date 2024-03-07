package dto

import (
	"errors"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	type scenario struct {
		name          string
		person        Person
		expectedError error
	}
	scenarios := []scenario{
		{
			name:          "empty nickname",
			person:        Person{},
			expectedError: errors.New("Apelido inválido"),
		},
		{
			name: "over len nickname",
			person: Person{
				Nickname: "asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas",
			},
			expectedError: errors.New("Apelido inválido"),
		},
		{
			name: "empty name",
			person: Person{
				Nickname: "Bob",
			},
			expectedError: errors.New("Nome inválido"),
		},
		{
			name: "over len name",
			person: Person{
				Nickname: "Bob",
				Name:     "asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas",
			},
			expectedError: errors.New("Nome inválido"),
		},
		{
			name: "invalid birth date",
			person: Person{
				Nickname:  "Bob",
				Name:      "Bob marley",
				BirthDate: "",
			},
			expectedError: errors.New("Data de nascimento inválida"),
		},
		{
			name: "invalid stack item",
			person: Person{
				Nickname:  "Bob",
				Name:      "Bob marley",
				BirthDate: "1945-02-06",
				Stacks:    []string{"asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas"},
			},
			expectedError: errors.New("Stack item inválido, asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas"),
		},
		{
			name: "valid person",
			person: Person{
				Nickname:  "Bob",
				Name:      "Bob marley",
				BirthDate: "1945-02-06",
				Stacks:    []string{"C#", "Golang", "Java"},
			},
			expectedError: nil,
		},
	}
	for _, scenario := range scenarios {
		err := scenario.person.Validate()
		require.True(t, reflect.DeepEqual(err, scenario.expectedError))
	}
}
