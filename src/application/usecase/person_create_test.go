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

func TestPersonCreateUseCaseExecute(t *testing.T) {
	type scenarions struct {
		name                                    string
		inputData                               interface{}
		outputError                             errors.CommonError
		outputCheckExistsNicknameServiceSuccess interface{}
		outputCheckExistsNicknameServiceError   errors.CommonError
		outputCreatePersonServiceSuccess        interface{}
		outputCreatePersonServiceError          errors.CommonError
	}

	tests := []scenarions{
		{
			name:        "invalid param intention",
			inputData:   entities.CreatePersonIntention{},
			outputError: errors.NewInvaliParamApiError("intention"),
		},
		{
			name:                                  "person already exists",
			inputData:                             &entities.CreatePersonIntention{Person: entities.Person{Nickname: "test"}},
			outputError:                           errors.NewPersonAlreadyExists(),
			outputCheckExistsNicknameServiceError: errors.NewPersonAlreadyExists(),
		},
		{
			name:                           "person create error in database",
			inputData:                      &entities.CreatePersonIntention{Person: entities.Person{Nickname: "test"}},
			outputError:                    errors.NewClientError(500, "person create error in database", ""),
			outputCreatePersonServiceError: errors.NewClientError(500, "person create error in database", ""),
		},
		{
			name:                             "person created",
			inputData:                        &entities.CreatePersonIntention{Person: entities.Person{Nickname: "test"}},
			outputError:                      nil,
			outputCreatePersonServiceSuccess: "123",
		},
	}

	for _, test := range tests {
		checkExistsNicknameServiceMock := service.NewCheckExistsNicknameServiceMock()
		checkExistsNicknameServiceMock.On("Execute", mock.Anything).Return(test.outputCheckExistsNicknameServiceSuccess, test.outputCheckExistsNicknameServiceError)

		personCreateServiceMock := service.NewPersonCreateServiceMock()
		personCreateServiceMock.On("Execute", mock.Anything).Return(test.outputCreatePersonServiceSuccess, test.outputCreatePersonServiceError)

		personCreateUseCase := NewPersonCreateUseCase(checkExistsNicknameServiceMock, personCreateServiceMock)

		t.Run(test.name, func(t *testing.T) {
			err := personCreateUseCase.Execute(test.inputData)

			require.True(t, reflect.DeepEqual(test.outputError, err))
		})
	}
}
