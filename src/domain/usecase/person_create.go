package usecase

import (
	inUsecase "docker-example/src/ports/in/usecase"
)

type personCreate struct {
}

func NewPersonCreate() inUsecase.PersonCreate {
	return &personCreate{}
}

func (p *personCreate) Execute() (err error) {
	return nil
}
