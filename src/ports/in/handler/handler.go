package handler

import (
	"docker-example/src/ports/in/drive/web/dto"
	"net/http"
)

type Handler interface {
	Ping(requestPathParam map[string][]string, requestQueryParam map[string][]string,
		requestHeaders map[string][]string, requestBody interface{}) (interface{}, error)
	PersonCreate(requestPathParam map[string][]string, requestQueryParam map[string][]string,
		requestHeaders map[string][]string, requestBody interface{}) (interface{}, error)
}

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (handler *handler) Ping(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody interface{}) (interface{}, error) {
	return &dto.DefaultResponse{
		ResponseBody: dto.ResponsePingDto{
			Message: "pong",
		},
		ResponseCode:    http.StatusOK,
		ResponseHeaders: map[string]string{},
	}, nil
}

func (handler *handler) PersonCreate(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody interface{}) (interface{}, error) {
	return &dto.DefaultResponse{
		ResponseBody: dto.ResponsePingDto{
			Message: "pong",
		},
		ResponseCode:    http.StatusOK,
		ResponseHeaders: map[string]string{},
	}, nil
}
