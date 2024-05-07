package dto

import (
	"docker-example/src/commons/errors"
	"reflect"
)

type RequestGetPersonByTerm struct {
	Term string
}

func (request RequestGetPersonByTerm) Validate() errors.CommonError {
	if len(request.Term) == 0 {
		return errors.NewInvaliFieldError("t")
	}
	return nil
}

func (person *RequestGetPersonByTerm) IsEmpty() bool {
	return reflect.DeepEqual(&RequestGetPersonByTerm{}, person)
}
