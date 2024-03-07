package handler

type ContractHandler func(requestPathParam map[string][]string, requestQueryParam map[string][]string, requestHeaders map[string][]string, requestBody interface{}) (interface{}, error)
