package usecase

import "docker-example/src/commons/errors"

type UseCase interface {
	Execute(intention interface{}) (err errors.CommonError)
}
