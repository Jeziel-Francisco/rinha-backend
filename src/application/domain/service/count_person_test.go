package service

import (
	"docker-example/src/commons/errors"
	"docker-example/src/ports/out/drive/adapter/person"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountPersonExecute(t *testing.T) {
	type scenarions struct {
		name                   string
		outputServiceSuccess   interface{}
		outputServiceError     errors.CommonError
		responseSuccessAdapter interface{}
		responseErrorAdapter   errors.CommonError
	}

	var quantityPerson uint64 = 123

	tests := []scenarions{
		{
			name:                 "integration database error",
			outputServiceError:   errors.NewClientError(500, "integration database error", ""),
			responseErrorAdapter: errors.NewClientError(500, "integration database error", ""),
		},
		{
			name:                   "count person success",
			outputServiceSuccess:   &quantityPerson,
			responseSuccessAdapter: &quantityPerson,
		},
	}

	for _, test := range tests {
		personAdapterMock := person.NewPersonDatabaseAdapterMock()
		personAdapterMock.On("CountPerson").Return(test.responseSuccessAdapter, test.responseErrorAdapter)

		countPersonService := NewCountPersonService(personAdapterMock)

		t.Run(test.name, func(t *testing.T) {
			var result interface{}
			var err errors.CommonError

			result, err = countPersonService.Execute()

			require.True(t, reflect.DeepEqual(test.outputServiceSuccess, result))

			require.True(t, reflect.DeepEqual(test.outputServiceError, err))
		})
	}
}
