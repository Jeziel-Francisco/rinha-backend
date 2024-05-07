package handler

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/commons/errors"
	defaultDto "docker-example/src/ports/in/drive/web/dto"
	"docker-example/src/ports/in/handler/dto"
	"docker-example/src/ports/in/usecase"
	"encoding/json"
	"net/http"
)

type Handler interface {
	Ping(requestPathParam map[string][]string, requestQueryParam map[string][]string,
		requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError)
	PersonCreate(requestPathParam map[string][]string, requestQueryParam map[string][]string,
		requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError)
	GetPersonByID(requestPathParam map[string][]string, requestQueryParam map[string][]string,
		requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError)
}

type handler struct {
	personCreateUseCase  usecase.UseCase
	getPersonByIDUseCase usecase.UseCase
}

func NewHandler(personCreateUseCase usecase.UseCase, getPersonByIDUseCase usecase.UseCase) Handler {
	return &handler{
		personCreateUseCase:  personCreateUseCase,
		getPersonByIDUseCase: getPersonByIDUseCase,
	}
}

func (handler *handler) Ping(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError) {
	return &defaultDto.DefaultResponse{
		ResponseBody: defaultDto.ResponsePingDto{
			Message: "pong",
		},
		ResponseCode:    http.StatusOK,
		ResponseHeaders: map[string]string{},
	}, nil
}

func (handler *handler) PersonCreate(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError) {

	var inputData dto.RequestCreatePersonDto
	if err := json.Unmarshal(requestBody, &inputData); err != nil {
		return nil, errors.NewUnmarshalApiError(err.Error())
	}

	if err := inputData.Validate(); err != nil {
		return nil, err
	}

	intention := &entities.CreatePersonIntention{
		Person: entities.Person{
			Nickname:  inputData.Nickname,
			Name:      inputData.Name,
			BirthDate: inputData.BirthDate,
			Stacks:    inputData.Stacks,
		},
	}

	if err := handler.personCreateUseCase.Execute(intention); err != nil {
		return nil, err
	}

	responseBody := &dto.ResponseCreatePersonDto{ID: intention.Person.ID}

	return &defaultDto.DefaultResponse{
		ResponseBody: responseBody,
		ResponseCode: http.StatusCreated,
		ResponseHeaders: map[string]string{
			"location": intention.Person.ID,
		},
	}, nil
}

func (handler *handler) GetPersonByID(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError) {

	inputData := dto.RequestGetPersonByID{
		ID: requestPathParam["id"][0],
	}

	if err := inputData.Validate(); err != nil {
		return nil, err
	}

	intention := &entities.GetPersonByIDIntention{
		Person: entities.Person{
			ID: inputData.ID,
		},
	}

	if err := handler.getPersonByIDUseCase.Execute(intention); err != nil {
		return nil, err
	}

	responseBody := &dto.ResponseGetPersonDetail{
		ID:        intention.Person.ID,
		Nickname:  intention.Person.Nickname,
		Name:      intention.Person.Name,
		BirthDate: intention.Person.BirthDate,
		Stacks:    intention.Person.Stacks,
	}

	return &defaultDto.DefaultResponse{
		ResponseBody:    responseBody,
		ResponseCode:    http.StatusOK,
		ResponseHeaders: map[string]string{},
	}, nil
}
