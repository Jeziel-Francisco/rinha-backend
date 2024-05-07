package mapper

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/ports/out/drive/database/dto"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestFromPersonToRequestCreatePersonDto(t *testing.T) {
	type scenarions struct {
		name      string
		inputData *entities.Person
		output    *dto.RequestCreatePersonDto
	}

	tests := []scenarions{
		{
			name:      "input data is nil",
			inputData: nil,
			output:    nil,
		},
		{
			name:      "input data is empty",
			inputData: &entities.Person{},
			output:    nil,
		},
		{
			name: "is valid input data",
			inputData: &entities.Person{
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
			output: &dto.RequestCreatePersonDto{
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FromPersonToRequestCreatePersonDto(test.inputData)

			require.True(t, reflect.DeepEqual(test.output, result))
		})
	}
}

func TestFromPersonDtoToPerson(t *testing.T) {
	type scenarions struct {
		name      string
		inputData *dto.ResponseGetPersonDto
		output    *entities.Person
	}

	personID := uuid.New()

	tests := []scenarions{
		{
			name:      "input data is nil",
			inputData: nil,
			output:    nil,
		},
		{
			name:      "input data is empty",
			inputData: &dto.ResponseGetPersonDto{},
			output:    nil,
		},
		{
			name: "is valid input data",
			inputData: &dto.ResponseGetPersonDto{
				ID:        personID,
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
			output: &entities.Person{
				ID:        personID.String(),
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FromResponseGetPersonDtoToPerson(test.inputData)

			require.True(t, reflect.DeepEqual(test.output, result))
		})
	}
}

func TestFromListPersonDtoToListPerson(t *testing.T) {
	type scenarions struct {
		name      string
		inputData []*dto.ResponseGetPersonDto
		output    []*entities.Person
	}

	personID := uuid.New()

	tests := []scenarions{
		{
			name:      "input data is nil",
			inputData: nil,
			output:    nil,
		},
		{
			name:      "input data is empty",
			inputData: []*dto.ResponseGetPersonDto{&dto.ResponseGetPersonDto{}},
			output:    nil,
		},
		{
			name: "is valid input data",
			inputData: []*dto.ResponseGetPersonDto{
				&dto.ResponseGetPersonDto{
					ID:        personID,
					Nickname:  "test",
					Name:      "test",
					BirthDate: "test",
					Stacks:    []string{"test1", "test2"},
				},
			},
			output: []*entities.Person{
				&entities.Person{
					ID:        personID.String(),
					Nickname:  "test",
					Name:      "test",
					BirthDate: "test",
					Stacks:    []string{"test1", "test2"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FromListResponseGetPersonDtoToListPerson(test.inputData)

			require.True(t, reflect.DeepEqual(test.output, result))
		})
	}
}
