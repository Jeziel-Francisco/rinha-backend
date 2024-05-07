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

func FromListResponseGetPersonDtoToListPerson(sources []*dto.ResponseGetPersonDto) []*entities.Person {
	if len(sources) == 0 {
		return nil
	}
	var result []*entities.Person
	for _, source := range sources {
		if source.IsEmpty() {
			continue
		}
		result = append(result, FromResponseGetPersonDtoToPerson(source))
	}
	return result
}
