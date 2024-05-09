package usecase

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/domain/service"
	"docker-example/src/commons/errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountPersonUseCaseExecute(t *testing.T) {
	type scenarions struct {
		name                            string
		inputData                       interface{}
		outputError                     errors.CommonError
		outputCountPersonServiceSuccess interface{}
		outputCountPersonServiceError   errors.CommonError
	}

	var quantityPerson uint64 = 123
	tests := []scenarions{
		{
			name:        "invalid param intention",
			inputData:   entities.CreatePersonIntention{},
			outputError: errors.NewInvaliParamApiError("intention"),
		},
		{
			name:                          "error in service",
			inputData:                     &entities.CountPersonIntention{},
			outputCountPersonServiceError: errors.NewInvaliParamApiError("args"),
			outputError:                   errors.NewInvaliParamApiError("args"),
		},
		{
			name:                            "count person success",
			inputData:                       &entities.CountPersonIntention{},
			outputCountPersonServiceSuccess: &quantityPerson,
		},
	}

	for _, test := range tests {
		countPersonServiceMock := service.NewCountPersonServiceMock()
		countPersonServiceMock.On("Execute").Return(test.outputCountPersonServiceSuccess, test.outputCountPersonServiceError)

		countPersonUseCase := NewCountPersonUseCase(countPersonServiceMock)

		t.Run(test.name, func(t *testing.T) {
			err := countPersonUseCase.Execute(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}
