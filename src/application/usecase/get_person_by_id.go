package usecase

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/domain/service"
	"docker-example/src/commons/errors"
	inUsecase "docker-example/src/ports/in/usecase"
)

type getPersonByIDUseCase struct {
	getPersonByIDService service.Service
}

func NewGetPersonByIDUseCase(getPersonByIDService service.Service) inUsecase.UseCase {
	return &getPersonByIDUseCase{
		getPersonByIDService: getPersonByIDService,
	}
}

func (useCase *getPersonByIDUseCase) Execute(intention interface{}) (err errors.CommonError) {
	getPersonByIDIntention, ok := intention.(*entities.GetPersonByIDIntention)
	if !ok {
		return errors.NewInvaliParamApiError("intention")
	}
	person, err := useCase.getPersonByIDService.Execute(getPersonByIDIntention.Person.ID)
	if err != nil {
		return err
	}

	getPersonByIDIntention.Person = *person.(*entities.Person)

	return nil
}
