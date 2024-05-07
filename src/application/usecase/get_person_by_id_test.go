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

func TestGetPersonByIDUseCaseExecute(t *testing.T) {
	type scenarions struct {
		name                              string
		inputData                         interface{}
		outputError                       errors.CommonError
		outputGetPersonByIDServiceSuccess interface{}
		outputGetPersonByIDServiceError   errors.CommonError
	}

	tests := []scenarions{
		{
			name:        "invalid param intention",
			inputData:   entities.CreatePersonIntention{},
			outputError: errors.NewInvaliParamApiError("intention"),
		},
		{
			name:                            "error in service",
			inputData:                       &entities.GetPersonByIDIntention{},
			outputGetPersonByIDServiceError: errors.NewInvaliParamApiError("args"),
			outputError:                     errors.NewInvaliParamApiError("args"),
		},
		{
			name:                              "person exists",
			inputData:                         &entities.GetPersonByIDIntention{},
			outputGetPersonByIDServiceSuccess: &entities.Person{ID: "123"},
		},
	}

	for _, test := range tests {
		getPersonByIDServiceMock := service.NewGetPersonByIDServiceMock()
		getPersonByIDServiceMock.On("Execute", mock.Anything).Return(test.outputGetPersonByIDServiceSuccess, test.outputGetPersonByIDServiceError)

		getPersonByIDUseCase := NewGetPersonByIDUseCase(getPersonByIDServiceMock)

		t.Run(test.name, func(t *testing.T) {
			err := getPersonByIDUseCase.Execute(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}
