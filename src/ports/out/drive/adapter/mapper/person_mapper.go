package mapper

import (
	"docker-example/src/application/domain/entities"
	"docker-example/src/ports/out/drive/database/dto"
)

func FromPersonToRequestCreatePersonDto(source *entities.Person) *dto.RequestCreatePersonDto {
	if source == nil || source.IsEmpty() {
		return nil
	}
	return &dto.RequestCreatePersonDto{
		Nickname:  source.Nickname,
		Name:      source.Name,
		BirthDate: source.BirthDate,
		Stacks:    source.Stacks,
	}
}

func FromResponseGetPersonDtoToPerson(source *dto.ResponseGetPersonDto) *entities.Person {
	if source == nil || source.IsEmpty() {
		return nil
	}
	return &entities.Person{
		ID:        source.ID.String(),
		Nickname:  source.Nickname,
		Name:      source.Name,
		BirthDate: source.BirthDate,
		Stacks:    source.Stacks,
	}
}
