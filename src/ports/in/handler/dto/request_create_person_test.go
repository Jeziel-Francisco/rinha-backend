package dto

import (
	"docker-example/src/commons/errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	type scenario struct {
		name                   string
		RequestCreatePersonDto RequestCreatePersonDto
		expectedError          errors.CommonError
	}
	scenarios := []scenario{
		{
			name:                   "empty nickname",
			RequestCreatePersonDto: RequestCreatePersonDto{},
			expectedError:          errors.NewInvaliFieldError("apelido"),
		},
		{
			name: "over len nickname",
			RequestCreatePersonDto: RequestCreatePersonDto{
				Nickname: "asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas",
			},
			expectedError: errors.NewInvaliFieldError("apelido"),
		},
		{
			name: "empty name",
			RequestCreatePersonDto: RequestCreatePersonDto{
				Nickname: "Bob",
			},
			expectedError: errors.NewInvaliFieldError("nome"),
		},
		{
			name: "over len name",
			RequestCreatePersonDto: RequestCreatePersonDto{
				Nickname: "Bob",
				Name:     "asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas",
			},
			expectedError: errors.NewInvaliFieldError("nome"),
		},
		{
			name: "invalid birth date",
			RequestCreatePersonDto: RequestCreatePersonDto{
				Nickname:  "Bob",
				Name:      "Bob marley",
				BirthDate: "",
			},
			expectedError: errors.NewInvaliFieldError("nascimento"),
		},
		{
			name: "invalid stack item",
			RequestCreatePersonDto: RequestCreatePersonDto{
				Nickname:  "Bob",
				Name:      "Bob marley",
				BirthDate: "1945-02-06",
				Stacks:    []string{"asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas"},
			},
			expectedError: errors.NewInvaliFieldError("asdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfasasdfgasfas"),
		},
		{
			name: "valid RequestCreatePersonDto",
			RequestCreatePersonDto: RequestCreatePersonDto{
				Nickname:  "Bob",
				Name:      "Bob marley",
				BirthDate: "1945-02-06",
				Stacks:    []string{"C#", "Golang", "Java"},
			},
			expectedError: nil,
		},
	}
	for _, scenario := range scenarios {
		err := scenario.RequestCreatePersonDto.Validate()
		require.True(t, reflect.DeepEqual(err, scenario.expectedError))
	}
}
