package person

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/database/dto"
	"docker-example/src/ports/out/drive/database/person"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetPersonByNickname(t *testing.T) {
	type scenarions struct {
		name                  string
		inputData             string
		outputError           errors.CommonError
		outputSuccess         *entities.Person
		outputDatabaseError   errors.CommonError
		outputDatabaseSuccess *dto.ResponseGetPersonDto
	}

	personID := uuid.New()

	tests := []scenarions{
		{
			name:                "error in get person",
			inputData:           "123",
			outputError:         errors.NewClientError(500, "error in database", ""),
			outputDatabaseError: errors.NewClientError(500, "error in database", ""),
		},
		{
			name:      "success get person",
			inputData: "123",
			outputSuccess: &entities.Person{
				ID:        personID.String(),
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
			outputDatabaseSuccess: &dto.ResponseGetPersonDto{
				ID:        personID,
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
		},
	}

	for _, test := range tests {
		personDatabaseMock := person.NewPersonDatabaseMock()
		personDatabaseMock.On("GetPersonByNickname", mock.Anything).Return(test.outputDatabaseSuccess, test.outputDatabaseError)

		personDatabaseAdapter := NewPersonDatabaseAdapter(personDatabaseMock)

		t.Run(test.name, func(t *testing.T) {
			result, err := personDatabaseAdapter.GetPersonByNickname(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputSuccess, result))
			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}

func TestGetPersonByID(t *testing.T) {
	type scenarions struct {
		name                  string
		inputData             string
		outputError           errors.CommonError
		outputSuccess         *entities.Person
		outputDatabaseError   errors.CommonError
		outputDatabaseSuccess *dto.ResponseGetPersonDto
	}

	personID := uuid.New()

	tests := []scenarions{
		{
			name:                "error in get person",
			inputData:           "123",
			outputError:         errors.NewClientError(500, "error in database", ""),
			outputDatabaseError: errors.NewClientError(500, "error in database", ""),
		},
		{
			name:      "success get person",
			inputData: "123",
			outputSuccess: &entities.Person{
				ID:        personID.String(),
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
			outputDatabaseSuccess: &dto.ResponseGetPersonDto{
				ID:        personID,
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
		},
	}

	for _, test := range tests {
		personDatabaseMock := person.NewPersonDatabaseMock()
		personDatabaseMock.On("GetPersonByID", mock.Anything).Return(test.outputDatabaseSuccess, test.outputDatabaseError)

		personDatabaseAdapter := NewPersonDatabaseAdapter(personDatabaseMock)

		t.Run(test.name, func(t *testing.T) {
			result, err := personDatabaseAdapter.GetPersonByID(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputSuccess, result))
			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}

func TestGetPersonByTerm(t *testing.T) {
	type scenarions struct {
		name                  string
		inputData             string
		outputError           errors.CommonError
		outputSuccess         []*entities.Person
		outputDatabaseError   errors.CommonError
		outputDatabaseSuccess []*dto.ResponseGetPersonDto
	}

	personID := uuid.New()

	tests := []scenarions{
		{
			name:                "error in get person",
			inputData:           "123",
			outputError:         errors.NewClientError(500, "error in database", ""),
			outputDatabaseError: errors.NewClientError(500, "error in database", ""),
		},
		{
			name:      "success get person",
			inputData: "123",
			outputSuccess: []*entities.Person{
				&entities.Person{
					ID:        personID.String(),
					Nickname:  "test",
					Name:      "test",
					BirthDate: "test",
					Stacks:    []string{"test1", "test2"},
				},
			},
			outputDatabaseSuccess: []*dto.ResponseGetPersonDto{
				&dto.ResponseGetPersonDto{
					ID:        personID,
					Nickname:  "test",
					Name:      "test",
					BirthDate: "test",
					Stacks:    []string{"test1", "test2"},
				},
			},
		},
	}

	for _, test := range tests {
		personDatabaseMock := person.NewPersonDatabaseMock()
		personDatabaseMock.On("GetPersonByTerm", mock.Anything).Return(test.outputDatabaseSuccess, test.outputDatabaseError)

		personDatabaseAdapter := NewPersonDatabaseAdapter(personDatabaseMock)

		t.Run(test.name, func(t *testing.T) {
			result, err := personDatabaseAdapter.GetPersonByTerm(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputSuccess, result))
			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}

func TestCountPerson(t *testing.T) {
	type scenarions struct {
		name                  string
		outputError           errors.CommonError
		outputSuccess         *uint64
		outputDatabaseError   errors.CommonError
		outputDatabaseSuccess *uint64
	}

	var quantityPerson uint64 = 123

	tests := []scenarions{
		{
			name:                "error in count person",
			outputError:         errors.NewClientError(500, "error in database", ""),
			outputDatabaseError: errors.NewClientError(500, "error in database", ""),
		},
		{
			name:                  "success count person",
			outputSuccess:         &quantityPerson,
			outputDatabaseSuccess: &quantityPerson,
		},
	}

	for _, test := range tests {
		personDatabaseMock := person.NewPersonDatabaseMock()
		personDatabaseMock.On("CountPerson", mock.Anything).Return(test.outputDatabaseSuccess, test.outputDatabaseError)

		personDatabaseAdapter := NewPersonDatabaseAdapter(personDatabaseMock)

		t.Run(test.name, func(t *testing.T) {
			result, err := personDatabaseAdapter.CountPerson()

			require.True(t, reflect.DeepEqual(test.outputSuccess, result))
			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}

func TestCreatePerson(t *testing.T) {
	type scenarions struct {
		name                  string
		inputData             *entities.Person
		outputError           errors.CommonError
		outputSuccess         *entities.Person
		outputDatabaseError   errors.CommonError
		outputDatabaseSuccess *dto.ResponseCreatePersonDto
	}

	personID := uuid.New()

	tests := []scenarions{
		{
			name: "error in create person",
			inputData: &entities.Person{
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
			outputError:         errors.NewClientError(500, "error in database", " "),
			outputDatabaseError: errors.NewClientError(500, "error in database", " "),
		},
		{
			name: "success create person",
			inputData: &entities.Person{
				Nickname:  "test",
				Name:      "test",
				BirthDate: "test",
				Stacks:    []string{"test1", "test2"},
			},
			outputSuccess: &entities.Person{
				ID: personID.String(),
			},
			outputDatabaseSuccess: &dto.ResponseCreatePersonDto{
				ID: personID,
			},
		},
	}

	for _, test := range tests {
		personDatabaseMock := person.NewPersonDatabaseMock()
		personDatabaseMock.On("Create", mock.Anything).Return(test.outputDatabaseSuccess, test.outputDatabaseError)

		personDatabaseAdapter := NewPersonDatabaseAdapter(personDatabaseMock)

		t.Run(test.name, func(t *testing.T) {
			result, err := personDatabaseAdapter.Create(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputSuccess, result))
			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}
