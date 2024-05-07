package handler

import (
	"docker-example/src/commons/errors"
	defaultDto "docker-example/src/ports/in/drive/web/dto"
	"docker-example/src/ports/in/handler/dto"
	"docker-example/src/ports/in/handler/mapper"
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
	GetPersonByTerm(requestPathParam map[string][]string, requestQueryParam map[string][]string,
		requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError)
}

type handler struct {
	personCreateUseCase    usecase.UseCase
	getPersonByIDUseCase   usecase.UseCase
	getPersonByTermUseCase usecase.UseCase
}

func NewHandler(personCreateUseCase usecase.UseCase, getPersonByIDUseCase usecase.UseCase, getPersonByTermUseCase usecase.UseCase) Handler {
	return &handler{
		personCreateUseCase:    personCreateUseCase,
		getPersonByIDUseCase:   getPersonByIDUseCase,
		getPersonByTermUseCase: getPersonByTermUseCase,
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

	intention := mapper.FromRequestCreatePersonDtoToCreatePersonIntention(&inputData)

	if err := handler.personCreateUseCase.Execute(intention); err != nil {
		return nil, err
	}

	return &defaultDto.DefaultResponse{
		ResponseBody: mapper.FromPersonIDToResponseCreatePersonDto(&intention.Person.ID),
		ResponseCode: http.StatusCreated,
		ResponseHeaders: map[string]string{
			"location": intention.Person.ID,
		},
	}, nil
}

func (handler *handler) GetPersonByID(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError) {

	inputData := &dto.RequestGetPersonByID{
		ID: requestPathParam["id"][0],
	}

	if err := inputData.Validate(); err != nil {
		return nil, err
	}

	intention := mapper.FromRequestGetPersonByIDToCreatePersonIntention(inputData)

	if err := handler.getPersonByIDUseCase.Execute(intention); err != nil {
		return nil, err
	}

	return &defaultDto.DefaultResponse{
		ResponseBody:    mapper.FromGetPersonByIDToResponseGetPersonDetail(&intention.Person),
		ResponseCode:    http.StatusOK,
		ResponseHeaders: map[string]string{},
	}, nil
}

func (handler *handler) GetPersonByTerm(requestPathParam map[string][]string, requestQueryParam map[string][]string,
	requestHeaders map[string][]string, requestBody []byte) (interface{}, errors.CommonError) {

	inputData := &dto.RequestGetPersonByTerm{
		Term: requestPathParam["t"][0],
	}

	if err := inputData.Validate(); err != nil {
		return nil, err
	}

	intention := mapper.FromRequestGetPersonByTermToCreatePersonIntention(inputData)

	if err := handler.getPersonByTermUseCase.Execute(intention); err != nil {
		return nil, err
	}

	return &defaultDto.DefaultResponse{
		ResponseBody:    mapper.FromGetPersonByTermToListResponseGetPersonDetail(intention.People),
		ResponseCode:    http.StatusOK,
		ResponseHeaders: map[string]string{},
	}, nil
}
