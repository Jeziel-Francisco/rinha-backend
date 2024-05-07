package usecase

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/domain/service"
	"docker-example/src/commons/errors"
	inUsecase "docker-example/src/ports/in/usecase"
)

type getPersonByTermUseCase struct {
	getPersonByTermService service.Service
}

func NewGetPersonByTermUseCase(getPersonByTermService service.Service) inUsecase.UseCase {
	return &getPersonByTermUseCase{
		getPersonByTermService: getPersonByTermService,
	}
}

func (useCase *getPersonByTermUseCase) Execute(intention interface{}) (err errors.CommonError) {
	getPersonByTermIntention, ok := intention.(*entities.GetPersonByTermIntention)
	if !ok {
		return errors.NewInvaliParamApiError("intention")
	}
	people, err := useCase.getPersonByTermService.Execute(getPersonByTermIntention.Term)
	if err != nil {
		return err
	}

	getPersonByTermIntention.People = people.([]*entities.Person)

	return nil
}
