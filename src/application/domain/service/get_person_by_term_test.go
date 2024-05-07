package service

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/adapter/person"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetPersonByTermExecute(t *testing.T) {
	type scenarions struct {
		name                   string
		inputData              interface{}
		outputServiceSuccess   interface{}
		outputServiceError     errors.CommonError
		responseSuccessAdapter interface{}
		responseErrorAdapter   errors.CommonError
	}

	tests := []scenarions{
		{
			name:               "not send params",
			outputServiceError: errors.NewInvaliParamApiError("args"),
		},
		{
			name:               "invalid params",
			inputData:          1,
			outputServiceError: errors.NewInvaliParamApiError("id"),
		},
		{
			name:                 "integration database error",
			inputData:            "123",
			outputServiceError:   errors.NewClientError(500, "integration database error", ""),
			responseErrorAdapter: errors.NewClientError(500, "integration database error", ""),
		},
		{
			name:      "person exsits",
			inputData: "123",
			outputServiceSuccess: []*entities.Person{
				{
					ID: "123",
				},
			},
			responseSuccessAdapter: []*entities.Person{
				{
					ID: "123",
				},
			},
		},
	}

	for _, test := range tests {
		personAdapterMock := person.NewPersonDatabaseAdapterMock()
		personAdapterMock.On("GetPersonByTerm", mock.Anything).Return(test.responseSuccessAdapter, test.responseErrorAdapter)

		getPersonByTermService := NewGetPersonByTermService(personAdapterMock)

		t.Run(test.name, func(t *testing.T) {
			var result interface{}
			var err errors.CommonError

			if test.inputData != nil {
				result, err = getPersonByTermService.Execute(test.inputData)
			} else {
				result, err = getPersonByTermService.Execute()
			}

			require.True(t, reflect.DeepEqual(test.outputServiceSuccess, result))

			require.True(t, reflect.DeepEqual(test.outputServiceError, err))
		})
	}
}
