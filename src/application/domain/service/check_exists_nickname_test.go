package service

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/ports/out/drive/adapter/person"

	"docker-example/src/commons/errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCheckExistsNicknameExecute(t *testing.T) {
	type scenarions struct {
		name                   string
		inputData              interface{}
		outputServiceSuccess   interface{}
		outputServiceError     errors.CommonError
		responseSuccessAdapter *entities.Person
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
			outputServiceError: errors.NewInvaliParamApiError("nickname"),
		},
		{
			name:                 "integration database error",
			inputData:            "test",
			outputServiceError:   errors.NewClientError(500, "integration database error", ""),
			responseErrorAdapter: errors.NewClientError(500, "integration database error", ""),
		},
		{
			name:                   "person already exists",
			inputData:              "test",
			outputServiceError:     errors.NewPersonAlreadyExists(),
			responseSuccessAdapter: &entities.Person{},
		},
		{
			name:      "person exsits",
			inputData: "test",
		},
	}

	for _, test := range tests {
		personAdapterMock := person.NewPersonDatabaseAdapterMock()
		personAdapterMock.On("GetPersonByNickname", mock.Anything).Return(test.responseSuccessAdapter, test.responseErrorAdapter)

		checkExistsNicknameService := NewCheckExistsNicknameService(personAdapterMock)

		t.Run(test.name, func(t *testing.T) {
			var result interface{}
			var err errors.CommonError

			if test.inputData != nil {
				result, err = checkExistsNicknameService.Execute(test.inputData)
			} else {
				result, err = checkExistsNicknameService.Execute()
			}

			require.True(t, reflect.DeepEqual(test.outputServiceSuccess, result))

			require.True(t, reflect.DeepEqual(test.outputServiceError, err))
		})
	}
}
