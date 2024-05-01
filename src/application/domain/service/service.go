package service

import "docker-example/src/commons/errors"

type Service interface {
	Execute(args ...interface{}) (interface{}, errors.CommonError)
}
