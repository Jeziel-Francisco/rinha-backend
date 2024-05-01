package usecase

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/application/domain/service"
	"docker-example/src/commons/errors"
	inUsecase "docker-example/src/ports/in/usecase"
)

type personCreateUseCase struct {
	checkExistsNickNameService service.Service
	personCreateService        service.Service
}

func NewPersonCreateUseCase(checkExistsNickNameService service.Service, personCreateService service.Service) inUsecase.UseCase {
	return &personCreateUseCase{
		checkExistsNickNameService: checkExistsNickNameService,
		personCreateService:        personCreateService,
	}
}

func (useCase *personCreateUseCase) Execute(intention interface{}) (err errors.CommonError) {
	createPersonIntention, ok := intention.(*entities.CreatePersonIntention)
	if !ok {
		return errors.NewInvaliParamApiError("intention")
	}
	_, err = useCase.checkExistsNickNameService.Execute(createPersonIntention.Person.Nickname)
	if err != nil {
		return err
	}
	ID, err := useCase.personCreateService.Execute(createPersonIntention.Person)
	if err != nil {
		return err
	}
	createPersonIntention.Person.ID = ID.(string)
	return nil
}
