package dto

import (
	"docker-example/src/commons/errors"

	"github.com/google/uuid"
)

type RequestGetPersonByID struct {
	ID string
}

func (request RequestGetPersonByID) Validate() errors.CommonError {
	if _, err := uuid.Parse(request.ID); err != nil {
		return errors.NewInvaliFieldError("id")
	}
	return nil
}
