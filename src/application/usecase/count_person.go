package usecase

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/domain/service"
	"docker-example/src/commons/errors"
	inUsecase "docker-example/src/ports/in/usecase"
)

type countPersonUseCase struct {
	countPersonService service.Service
}

func NewCountPersonUseCase(countPersonService service.Service) inUsecase.UseCase {
	return &countPersonUseCase{
		countPersonService: countPersonService,
	}
}

func (useCase *countPersonUseCase) Execute(intention interface{}) (err errors.CommonError) {
	countPersonIntention, ok := intention.(*entities.CountPersonIntention)
	if !ok {
		return errors.NewInvaliParamApiError("intention")
	}
	quantityPerson, err := useCase.countPersonService.Execute()
	if err != nil {
		return err
	}

	countPersonIntention.QuantityPerson = *quantityPerson.(*uint64)

	return nil
}
