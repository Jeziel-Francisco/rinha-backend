package usecase

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/domain/service"
	"docker-example/src/commons/errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetPersonByTermUseCaseExecute(t *testing.T) {
	type scenarions struct {
		name                                string
		inputData                           interface{}
		outputError                         errors.CommonError
		outputGetPersonByTermServiceSuccess interface{}
		outputGetPersonByTermServiceError   errors.CommonError
	}

	tests := []scenarions{
		{
			name:        "invalid param intention",
			inputData:   entities.CreatePersonIntention{},
			outputError: errors.NewInvaliParamApiError("intention"),
		},
		{
			name:                              "error in service",
			inputData:                         &entities.GetPersonByTermIntention{},
			outputGetPersonByTermServiceError: errors.NewInvaliParamApiError("args"),
			outputError:                       errors.NewInvaliParamApiError("args"),
		},
		{
			name:                                "person exists",
			inputData:                           &entities.GetPersonByTermIntention{Term: "123"},
			outputGetPersonByTermServiceSuccess: []*entities.Person{{ID: "123"}},
		},
	}

	for _, test := range tests {
		getPersonByTermServiceMock := service.NewGetPersonByTermServiceMock()
		getPersonByTermServiceMock.On("Execute", mock.Anything).Return(test.outputGetPersonByTermServiceSuccess, test.outputGetPersonByTermServiceError)

		getPersonByTermUseCase := NewGetPersonByTermUseCase(getPersonByTermServiceMock)

		t.Run(test.name, func(t *testing.T) {
			err := getPersonByTermUseCase.Execute(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}
