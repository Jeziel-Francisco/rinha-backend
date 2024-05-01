package handler

import "docker-example/src/commons/errors"

type ContractHandler func(requestPathParam map[string][]string, requestQueryParam map[string][]string, requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError)
